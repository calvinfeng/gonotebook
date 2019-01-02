package model

import "github.com/jinzhu/gorm"

// Message is a model for message entity.
type Message struct {
	gorm.Model
	UserID int    `gorm:"type:integer; index"` // foreign key, belongs to
	Body   string `gorm:"type:text"`
}
