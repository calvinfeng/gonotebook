package model

import "time"

// Message is a model for messages.
type Message struct {
	ID        uint      `gorm:"column:id"         json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
	Body      string    `gorm:"column:body"       json:"body"`
	UserID    uint      `gorm:"column:user_id"    json:"user_id"` // Foreign key
}

// TableName tells GORM where to find this record.
func (Message) TableName() string {
	return "messages"
}

// Validate performs validation on message model.
func (m *Message) Validate() error {
	if len(m.Body) == 0 {
		return &ValidationError{Field: "body", Message: "cannot be empty"}
	}

	if m.UserID == 0 {
		return &ValidationError{Field: "user_id", Message: "is required"}
	}

	return nil
}
