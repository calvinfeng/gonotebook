package handler

import (
	"github.com/calvinfeng/go-academy/userauth/model"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var log = logrus.WithFields(logrus.Fields{
	"pkg": "handler",
})

func findUserByToken(db *gorm.DB, token string) (*model.User, error) {
	var user model.User
	if err := db.Where("jwt_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func findUserByCredentials(db *gorm.DB, c *Credential) (*model.User, error) {
	var user model.User
	if err := db.Where("email = ?", c.Email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordDigest, []byte(c.Password)); err != nil {
		return nil, err
	}

	return &user, nil
}
