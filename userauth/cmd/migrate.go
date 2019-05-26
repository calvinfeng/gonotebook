package cmd

import (
	"fmt"

	"github.com/calvinfeng/go-academy/userauth/model"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres" // Driver
	_ "github.com/golang-migrate/migrate/source/file"       // Driver
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // Driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

const (
	host         = "localhost"
	port         = "5432"
	user         = "cfeng"
	password     = "cfeng"
	database     = "go_user_auth"
	ssl          = "sslmode=disable"
	migrationDir = "file://./migrations/"
)

var log = logrus.WithFields(logrus.Fields{
	"pkg": "cmd",
})

var pgAddress = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s", user, password, host, port, database, ssl)

// RunMigrationsCmd is a command to run migration.
var RunMigrationsCmd = &cobra.Command{
	Use:   "runmigrations",
	Short: "run migration on database",
	RunE:  runmigrations,
}

func runmigrations(cmd *cobra.Command, args []string) error {
	migration, err := migrate.New(migrationDir, pgAddress)
	if err != nil {
		return err
	}

	log.Info("performing reset on database")
	if err = migration.Drop(); err != nil {
		return err
	}

	if err := migration.Up(); err != nil {
		return err
	}

	log.Info("migration has been performed successfully")

	db, err := gorm.Open("postgres", pgAddress)
	if err != nil {
		return err
	}

	admin := &model.User{
		Name:     "Calvin Feng",
		Email:    "cfeng@goacademy.com",
		Password: "cfeng",
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return err
	}

	admin.PasswordDigest = hashBytes
	admin.JWTToken = "admin"

	if err := db.Create(admin).Error; err != nil {
		return err
	}

	log.Info("admin is created")
	return nil
}
