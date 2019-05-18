package cmd

import (
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres" // Driver
	_ "github.com/golang-migrate/migrate/source/file"       // Driver
	_ "github.com/lib/pq"                                   // Driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	up       = "up"
	reset    = "reset"
	host     = "localhost"
	port     = "5432"
	user     = "cfeng"
	password = "cfeng"
	database = "go_user_auth"
	ssl      = "sslmode=disable"
)

const migrationUsage = `
Commands:
	up                   Migrate the DB to the most recent version available
	reset                Resets the database
Usage:
	userauth migrate <command>
`

const migrationDir = "file://./migrations/"

// RunMigrationCmd is a command to run migration.
var RunMigrationCmd = &cobra.Command{
	Use:   "runmigration",
	Short: "run migration on database",
	RunE:  runmigration,
}

func runmigration(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		fmt.Println(migrationUsage)
		return fmt.Errorf("no commands provided")
	}

	addr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s", user, password, host, port, database, ssl)

	migration, err := migrate.New(migrationDir, addr)
	if err != nil {
		return err
	}

	command := args[0]
	switch command {
	case up:
		if err := migration.Up(); err != nil {
			return err
		}

		logrus.Info("migration has been performed successfully")
	case reset:
		if err = migration.Drop(); err != nil {
			return err
		}

		logrus.Info("database has been reset")
	}

	return nil
}
