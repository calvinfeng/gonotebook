package model

import "time"

// Message is a model for messages.
type Message struct {
	ID        uint      `gorm:"column:id"          json:"id"`
	CreatedAt time.Time `gorm:"column:created_at"  json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"  json:"-"`
	Body      string    `gorm:"column:body"        json:"body"`

	// Foreign keys
	SenderID   uint  `gorm:"column:sender_id"       json:"sender_id"`
	Sender     *User `gorm:"foreignkey:sender_id"   json:"sender,omitempty"`
	ReceiverID uint  `gorm:"column:receiver_id"     json:"receiver_id"`
	Receiver   *User `gorm:"foreignkey:receiver_id" json:"receiver,omitempty"`
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

	if m.SenderID == 0 {
		return &ValidationError{Field: "sender_id", Message: "required"}
	}

	if m.ReceiverID == 0 {
		return &ValidationError{Field: "receiver_id", Message: "required"}
	}

	return nil
}
