package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "userauth",
	Short: "user authentication service",
}

var runmigrationCmd = &cobra.Command{
	Use:   "runmigration",
	Short: "run migration on database",
	RunE:  runmigration,
}

var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "run user authentication server",
	RunE:  runserver,
}

// Execute configures command and executes them.
func Execute() {
	// Run it!
	rootCmd.AddCommand(runserverCmd, runmigrationCmd)
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}
