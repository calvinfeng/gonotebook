package cmd

import (
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres" // Driver
	_ "github.com/golang-migrate/migrate/source/file"       // Driver
	_ "github.com/lib/pq"                                   // Driver
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
	sslMode  = "sslmode=disable"
)

const migrationUsage = `
Commands:
	up                   Migrate the DB to the most recent version available
	reset                Resets the database
Usage:
	userauth migrate <command>
`

const migrationDir = "file://./migrations/"

func runmigration(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		fmt.Println(migrationUsage)
		return fmt.Errorf("no commands provided")
	}

	psqlAddress := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s",
		user, password, host, port, database, sslMode)

	migration, err := migrate.New(migrationDir, psqlAddress)
	if err != nil {
		return err
	}

	command := args[0]
	switch command {
	case up:
		if err := migration.Up(); err != nil {
			return err
		}
	case reset:
		if err = migration.Drop(); err != nil {
			return err
		}
	}

	return nil
}
