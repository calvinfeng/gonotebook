package model

import "github.com/jinzhu/gorm"

// User is a model for user entity.
type User struct {
	gorm.Model
	Name           string    `gorm:"type:varchar(100)"              json:"name"`
	Email          string    `gorm:"type:varchar(100);unique_index" json:"email"`
	SessionToken   string    `gorm:"type:varchar(100);unique_index" json:"-"`
	PasswordDigest []byte    `gorm:"type:bytea;unique_index"        json:"-"`
	Messages       []Message `gorm:"ForeignKey:UserID"              json:"-"` // has-many
}

// ResetSessionToken resets the token that user has.
func (u *User) ResetSessionToken() {
	if randStr, err := generateRandomString(20); err == nil {
		u.SessionToken = randStr
	}
}
