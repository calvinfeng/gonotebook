package handler

import (
	"encoding/json"
	"net/http"

	"github.com/calvinfeng/go-academy/userauth/model"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type (
	// ErrorResponse is a payload that returns to the client indicating there's error.
	ErrorResponse struct {
		Error string `json:"error"`
	}

	// RegisterRequest is a payload that client submits to register as a user.
	RegisterRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// UserJSONResponse is a payload that server returns to client upon successful registration.
	UserJSONResponse struct {
		Name         string   `json:"name"`
		Email        string   `json:"email"`
		SessionToken string   `json:"session_token"`
		Messages     []string `json:"messages"`
	}

	// LoginRequest is a payload that client submits.
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// LogoutResponse is a payload that server returns to client.
	LogoutResponse struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		IsLoggedOut bool   `json:"is_logged_out"`
	}
)

func renderError(w http.ResponseWriter, code int, msg string) {
	res := &ErrorResponse{msg}
	bytes, _ := json.Marshal(res)
	w.WriteHeader(code)
	w.Write(bytes)
}

func findUserByToken(db *gorm.DB, token string) (*model.User, error) {
	var user model.User
	if err := db.Where("session_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func findUserByCredential(db *gorm.DB, email, password string) (*model.User, error) {
	var user model.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordDigest, []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil
}
