package main

import (
	"os"

	"github.com/calvinfeng/go-academy/auctionhouse/cmd"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	root := &cobra.Command{
		Use:   "auctionhouse",
		Short: "an auction house running on gRPC and HTTP",
	}

	root.AddCommand(cmd.RunMigrationsCmd, cmd.RunServerCmd, cmd.RunClientCmd)
	if err := root.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
