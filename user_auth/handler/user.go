package handler

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"go-academy/user_auth/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserJSONResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	SessionToken string `json:"session_token"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserListHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []model.User

		if err := db.Find(&users).Error; err != nil {
			RenderError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := []*UserJSONResponse{}
		for _, user := range users {
			res = append(res, &UserJSONResponse{
				Name:         user.Name,
				Email:        user.Email,
				SessionToken: user.SessionToken,
			})
		}

		if bytes, err := json.Marshal(res); err != nil {
			RenderError(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}
	}
}

func NewUserCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var regReq RegisterRequest
		if err := decoder.Decode(&regReq); err != nil {
			RenderError(w, "Failed to parse request JSON into struct", http.StatusInternalServerError)
			return
		}

		if len(regReq.Email) == 0 || len(regReq.Password) == 0 || len(regReq.Name) == 0 {
			RenderError(w, "Please provide name, email and password for registration", http.StatusBadRequest)
			return
		}

		hashBytes, hashErr := bcrypt.GenerateFromPassword([]byte(regReq.Password), 10)
		if hashErr != nil {
			RenderError(w, hashErr.Error(), http.StatusInternalServerError)
			return
		}

		newUser := &model.User{
			Name:           regReq.Name,
			Email:          regReq.Email,
			PasswordDigest: hashBytes,
		}

		newUser.ResetSessionToken()

		if err := db.Create(&newUser).Error; err != nil {
			RenderError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Token is set to expire in 2 days
		expiration := time.Now().Add(2 * 24 * time.Hour)
		cookie := http.Cookie{Name: "session_token", Value: newUser.SessionToken, Expires: expiration}
		http.SetCookie(w, &cookie)

		res := &UserJSONResponse{
			Name:         newUser.Name,
			Email:        newUser.Email,
			SessionToken: newUser.SessionToken,
		}

		if bytes, err := json.Marshal(res); err != nil {
			RenderError(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}
	}
}
