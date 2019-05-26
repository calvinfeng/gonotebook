package main

import (
	"os"

	"github.com/calvinfeng/go-academy/userauth/cmd"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(logrus.DebugLevel)

	root := &cobra.Command{
		Use:   "userauth",
		Short: "user authentication service",
	}

	root.AddCommand(cmd.RunMigrationsCmd, cmd.RunServerCmd)
	if err := root.Execute(); err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}
