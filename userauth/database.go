package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-academy/user_auth/model"
)

// SetupDatabase will perform database connection and auto migration on all gorm.Models
func SetupDatabase() (*gorm.DB, error) {
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
