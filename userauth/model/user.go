package model

import "github.com/jinzhu/gorm"

// User is a model for user entity.
type User struct {
	gorm.Model
	Name           string    `gorm:"column:name"`
	Email          string    `gorm:"column:email"`
	SessionToken   string    `gorm:"column:session_token"`
	PasswordDigest []byte    `gorm:"column:password_digest"`
	Messages       []Message `gorm:"foreignkey:UserID"` // has-many
}

// ResetSessionToken resets the token that user has.
func (u *User) ResetSessionToken() {
	if randStr, err := generateRandomString(20); err == nil {
		u.SessionToken = randStr
	}
}
