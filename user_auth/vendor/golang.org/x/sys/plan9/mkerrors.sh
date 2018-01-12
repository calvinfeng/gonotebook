#!/usr/bin/env bash
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# Generate Go code listing errors and other #defined constant
# values (ENAMETOOLONG etc.), by asking the preprocessor
# about the definitions.

unset LANG
export LC_ALL=C
export LC_CTYPE=C

CC=${CC:-gcc}

uname=$(uname)

includes='
#include <sys/types.h>
#include <sys/file.h>
#include <fcntl.h>
#include <dirent.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netinet/ip.h>
#include <netinet/ip6.h>
#include <netinet/tcp.h>
#include <errno.h>
#include <sys/signal.h>
#include <signal.h>
#include <sys/resource.h>
'

ccflags="$@"

# Write go tool cgo -godefs input.
(
	echo package plan9
	echo
	echo '/*'
	indirect="includes_$(uname)"
	echo "${!indirect} $includes"
	echo '*/'
	echo 'import "C"'
	echo
	echo 'const ('

	# The gcc command line prints all the #defines
	# it encounters while processing the input
	echo "${!indirect} $includes" | $CC -X c - -E -dM $ccflags |
	awk '
		$1 != "#define" || $2 ~ /\(/ || $3 == "" {next}

		$2 ~ /^E([ABCD]X|[BIS]P|[SD]I|S|FL)$/ {next}  # 386 registers
		$2 ~ /^(SIGEV_|SIGSTKSZ|SIGRT(MIN|MAX))/ {next}
		$2 ~ /^(SCM_SRCRT)$/ {next}
		$2 ~ /^(MAP_FAILED)$/ {next}

		$2 !~ /^ETH_/ &&
		$2 !~ /^EPROC_/ &&
		$2 !~ /^EQUIV_/ &&
		$2 !~ /^EXPR_/ &&
		$2 ~ /^E[A-Z0-9_]+$/ ||
		$2 ~ /^B[0-9_]+$/ ||
		$2 ~ /^V[A-Z0-9]+$/ ||
		$2 ~ /^CS[A-Z0-9]/ ||
		$2 ~ /^I(SIG|CANON|CRNL|EXTEN|MAXBEL|STRIP|UTF8)$/ ||
		$2 ~ /^IGN/ ||
		$2 ~ /^IX(ON|ANY|OFF)$/ ||
		$2 ~ /^IN(LCR|PCK)$/ ||
		$2 ~ /(^FLU?SH)|(FLU?SH$)/ ||
		$2 ~ /^C(LOCAL|READ)$/ ||
		$2 == "BRKINT" ||
		$2 == "HUPCL" ||
		$2 == "PENDIN" ||
		$2 == "TOSTOP" ||
		$2 ~ /^PAR/ ||
		$2 ~ /^SIG[^_]/ ||
		$2 ~ /^O[CNPFP][A-Z]+[^_][A-Z]+$/ ||
		$2 ~ /^IN_/ ||
		$2 ~ /^LOCK_(SH|EX|NB|UN)$/ ||
		$2 ~ /^(AF|SOCK|SO|SOL|IPPROTO|IP|IPV6|ICMP6|TCP|EVFILT|NOTE|EV|SHUT|PROT|MAP|PACKET|MSG|SCM|MCL|DT|MADV|PR)_/ ||
		$2 == "ICMPV6_FILTER" ||
		$2 == "SOMAXCONN" ||
		$2 == "NAME_MAX" ||
		$2 == "IFNAMSIZ" ||
		$2 ~ /^CTL_(MAXNAME|NET|QUERY)$/ ||
		$2 ~ /^SYSCTL_VERS/ ||
		$2 ~ /^(MS|MNT)_/ ||
		$2 ~ /^TUN(SET|GET|ATTACH|DETACH)/ ||
		$2 ~ /^(O|F|FD|NAME|S|PTRACE|PT)_/ ||
		$2 ~ /^LINUX_REBOOT_CMD_/ ||
		$2 ~ /^LINUX_REBOOT_MAGIC[12]$/ ||
		$2 !~ "NLA_TYPE_MASK" &&
		$2 ~ /^(NETLINK|NLM|NLMSG|NLA|IFA|IFAN|RT|RTCF|RTN|RTPROT|RTNH|ARPHRD|ETH_P)_/ ||
		$2 ~ /^SIOC/ ||
		$2 ~ /^TIOC/ ||
		$2 !~ "RTF_BITS" &&
		$2 ~ /^(IFF|IFT|NET_RT|RTM|RTF|RTV|RTA|RTAX)_/ ||
		$2 ~ /^BIOC/ ||
		$2 ~ /^RUSAGE_(SELF|CHILDREN|THREAD)/ ||
		$2 ~ /^RLIMIT_(AS|CORE|CPU|DATA|FSIZE|NOFILE|STACK)|RLIM_INFINITY/ ||
		$2 ~ /^PRIO_(PROCESS|PGRP|USER)/ ||
		$2 ~ /^CLONE_[A-Z_]+/ ||
		$2 !~ /^(BPF_TIMEVAL)$/ &&
		$2 ~ /^(BPF|DLT)_/ ||
		$2 !~ "WMESGLEN" &&
		$2 ~ /^W[A-Z0-9]+$/ {printf("\t%s = C.%s\n", $2, $2)}
		$2 ~ /^__WCOREFLAG$/ {next}
		$2 ~ /^__W[A-Z0-9]+$/ {printf("\t%s = C.%s\n", substr($2,3), $2)}

		{next}
	' | sort

	echo ')'
) >_const.go

# Pull out the error names for later.
errors=$(
	echo '#include <errno.h>' | $CC -X c - -E -dM $ccflags |
	awk '$1=="#define" && $2 ~ /^E[A-Z0-9_]+$/ { print $2 }' |
	sort
)

# Pull out the signal names for later.
signals=$(
	echo '#include <signal.h>' | $CC -X c - -E -dM $ccflags |
	awk '$1=="#define" && $2 ~ /^SIG[A-Z0-9]+$/ { print $2 }' |
	egrep -v '(SIGSTKSIZE|SIGSTKSZ|SIGRT)' |
	sort
)

# Again, writing regexps to a file.
echo '#include <errno.h>' | $CC -X c - -E -dM $ccflags |
	awk '$1=="#define" && $2 ~ /^E[A-Z0-9_]+$/ { print "^\t" $2 "[ \t]*=" }' |
	sort >_error.grep
echo '#include <signal.h>' | $CC -X c - -E -dM $ccflags |
	awk '$1=="#define" && $2 ~ /^SIG[A-Z0-9]+$/ { print "^\t" $2 "[ \t]*=" }' |
	egrep -v '(SIGSTKSIZE|SIGSTKSZ|SIGRT)' |
	sort >_signal.grep

echo '// mkerrors.sh' "$@"
echo '// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT'
echo
go tool cgo -godefs -- "$@" _const.go >_error.out
cat _error.out | grep -vf _error.grep | grep -vf _signal.grep
echo
echo '// Errors'
echo 'const ('
cat _error.out | grep -f _error.grep | sed 's/=\(.*\)/= Errno(\1)/'
echo ')'

echo
echo '// Signals'
echo 'const ('
cat _error.out | grep -f _signal.grep | sed 's/=\(.*\)/= Signal(\1)/'
echo ')'

# Run C program to print error and syscall strings.
(
	echo -E "
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <ctype.h>
#include <string.h>
#include <signal.h>

#define nelem(X) (sizeof(X)/sizeof((X)[0]))

enum { A = 'A', Z = 'Z', a = 'a', z = 'z' }; // avoid need for single quotes below

int errors[] = {
"
	for i in $errors
	do
		echo -E '	'$i,
	done

	echo -E "
};

int signals[] = {
"
	for i in $signals
	do
		echo -E '	'$i,
	done

	# Use -E because on some systems bash builtin interprets \n itself.
	echo -E '
};

static int
intcmp(const void *a, const void *b)
{
	return *(int*)a - *(int*)b;
}

int
main(void)
{
	int i, j, e;
	char buf[1024], *p;

	printf("\n\n// Error table\n");
	printf("var errors = [...]string {\n");
	qsort(errors, nelem(errors), sizeof errors[0], intcmp);
	for(i=0; i<nelem(errors); i++) {
		e = errors[i];
		if(i > 0 && errors[i-1] == e)
			continue;
		strcpy(buf, strerror(e));
		// lowercase first letter: Bad -> bad, but STREAM -> STREAM.
		if(A <= buf[0] && buf[0] <= Z && a <= buf[1] && buf[1] <= z)
			buf[0] += a - A;
		printf("\t%d: \"%s\",\n", e, buf);
	}
	printf("}\n\n");
	
	printf("\n\n// Signal table\n");
	printf("var signals = [...]string {\n");
	qsort(signals, nelem(signals), sizeof signals[0], intcmp);
	for(i=0; i<nelem(signals); i++) {
		e = signals[i];
		if(i > 0 && signals[i-1] == e)
			continue;
		strcpy(buf, strsignal(e));
		// lowercase first letter: Bad -> bad, but STREAM -> STREAM.
		if(A <= buf[0] && buf[0] <= Z && a <= buf[1] && buf[1] <= z)
			buf[0] += a - A;
		// cut trailing : number.
		p = strrchr(buf, ":"[0]);
		if(p)
			*p = '\0';
		printf("\t%d: \"%s\",\n", e, buf);
	}
	printf("}\n\n");

	return 0;
}

'
) >_errors.c

$CC $ccflags -o _errors _errors.c && $GORUN ./_errors && rm -f _errors.c _errors _const.go _error.grep _signal.grep _error.out
