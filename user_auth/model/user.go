package model

import "github.com/jinzhu/gorm"

// Default table name is users
type User struct {
	gorm.Model
	Name           string    `gorm:"type:varchar(100)" json:"name"`
	Email          string    `gorm:"type:varchar(100);unique_index" json:"email"`
	SessionToken   string    `gorm:"type:varchar(100);unique_index"`
	PasswordDigest []byte    `gorm:"type:bytea;unique_index"`
	Messages       []Message `gorm:"ForeignKey:UserID"` // has-many
}

func (u *User) ResetSessionToken() {
	if randStr, err := GenerateRandomString(20); err == nil {
		u.SessionToken = randStr
	}
}
