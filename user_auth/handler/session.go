package handler

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"go-academy/user_auth/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type SessionJSONResponse struct {
	Email        string `json:"email"`
	SessionToken string `json:"session_token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewSessionCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var loginReq LoginRequest
		if err := decoder.Decode(&loginReq); err != nil {
			RenderError(w, "Failed to parse request JSON into struct", http.StatusInternalServerError)
			return
		}

		var user model.User
		if err := db.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
			RenderError(w, "Wrong email", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword(user.PasswordDigest, []byte(loginReq.Password)); err != nil {
			RenderError(w, "Incorrect password", http.StatusUnauthorized)
			return
		}

		res := SessionJSONResponse{
			Email:        user.Email,
			SessionToken: user.SessionToken,
		}

		if bytes, err := json.Marshal(res); err != nil {
			RenderError(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}
	}
}
