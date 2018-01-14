# Setup
## Dependency Management with `dep`
First of all, get `dep` for dependency management

```
go get -u github.com/golang/dep/cmd/dep
```

If you are using Mac OS Xtr then it's even easier

```
brew install dep 
brew upgrade dep
```

I love Homebrew on Mac. It has everything!

## Databse
I am going to use PostgreSQL for this project, so let's create one. The superuser on my computer is `cfeng` so I will use 
that to create a database named `go_user_auth`

If you don't have a role or wish to create a separate role for this project, then just do the following
```
$ psql postgres
postgres=# create role <name> superuser login;
```

Create a database named `go_user_auth` with owner pointing to whichever account you like. I am using cfeng on my computer.
```
$ psql postgres
postgres=# create database go_user_auth with owner=cfeng;
```


Then just exit with `\q`

Actually just in case you don't remember the password to your `ROLE`, do the following
```
postgres=# alter user <your_username> with password <whatever you like>
```

I did mine with
```
postgres=# alter user cfeng with password "cfeng";
```



