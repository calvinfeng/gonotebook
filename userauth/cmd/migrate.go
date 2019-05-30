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
	database     = "go_academy_userauth"
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

func createUserList(db *gorm.DB, users []*model.User) error {
	for i, u := range users {
		hashBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
		if err != nil {
			return err
		}

		u.PasswordDigest = hashBytes
		u.JWTToken = fmt.Sprintf("admin-%d", i)
		if err := db.Create(u).Error; err != nil {
			return err
		}
		logrus.Infof("user %s is created", u.Email)
	}
	return nil
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

	return createUserList(db, []*model.User{
		&model.User{
			Name:     "Alice",
			Email:    "alice@goacademy.com",
			Password: "alice",
		},
		&model.User{
			Name:     "Bob",
			Email:    "bob@goacademy.com",
			Password: "bob",
		},
		&model.User{
			Name:     "Calvin",
			Email:    "calvin@goacademy.com",
			Password: "calvin",
		},
	})
}
