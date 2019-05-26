package model

import "time"

// User is a user model.
type User struct {
	// Database only
	ID             uint      `gorm:"column:id"              json:"-"`
	CreatedAt      time.Time `gorm:"column:created_at"      json:"-"`
	UpdatedAt      time.Time `gorm:"column:updated_at"      json:"-"`
	PasswordDigest []byte    `gorm:"column:password_digest" json:"-"`
	JWTToken       string    `gorm:"column:jwt_token"       json:"-"`

	// JSON only
	Password string `sql:"-" json:"password,omitempty"`

	// Both
	Name     string     `gorm:"column:name"        json:"name" `
	Email    string     `gorm:"column:email"       json:"email"`
	Messages []*Message `gorm:"foreignkey:user_id" json:"messages,omitempty"`
}

// TableName tells GORM where to find this record.
func (User) TableName() string {
	return "users"
}

// Validate performs validation on user model.
func (u *User) Validate() error {
	if len(u.Name) == 0 {
		return &ValidationError{Field: "name", Message: "cannot be empty"}
	}

	if len(u.Email) == 0 {
		return &ValidationError{Field: "email", Message: "cannot be empty"}
	}

	if len(u.Password) == 0 {
		return &ValidationError{Field: "password", Message: "too short"}
	}

	return nil
}

// ResetJWTToken resets the fake JWT token.
func (u *User) ResetJWTToken() {
	if randStr, err := fakeJWTToken(20); err == nil {
		u.JWTToken = randStr
	}
}
