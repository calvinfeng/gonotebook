package main

import (
	"go-academy/userauth/model"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

const addr = ":3000"

func connectDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		"postgres",
		"user=cfeng password=cfeng dbname=go_user_auth sslmode=disable",
	)

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})

	return db, nil
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	db, err := connectDB()

	if err != nil {
		logrus.Error(err)
		return
	}

	defer db.Close()

	server := &http.Server{
		Handler:      loadRoutes(db),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Infof("HTTP server is listening and serving on port %v", addr)
	logrus.Fatal(server.ListenAndServe())
}
