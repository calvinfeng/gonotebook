package model

import "github.com/jinzhu/gorm"

// Message is a model for message entity.
type Message struct {
	gorm.Model
	UserID int    `gorm:"column:user_id"` // Foreign key, belongs to
	Body   string `gorm:"column:body"`
}
