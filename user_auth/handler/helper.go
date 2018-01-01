package handler

import (
	"github.com/jinzhu/gorm"
	"go-academy/user_auth/model"
	"golang.org/x/crypto/bcrypt"
)

func FindUserByToken(db *gorm.DB, token string) (*model.User, error) {
	var user model.User
	if err := db.Where("session_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUserByCredential(db *gorm.DB, email, password string) (*model.User, error) {
	var user model.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordDigest, []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil
}
