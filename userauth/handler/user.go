package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/calvinfeng/go-academy/userauth/model"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// NewUserListHandler returns a handler that renders the list of users on the server.
func NewUserListHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []model.User

		if err := db.Preload("Messages").Find(&users).Error; err != nil {
			renderError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := []*UserJSONResponse{}
		for _, user := range users {
			msgs := []string{}
			for _, msg := range user.Messages {
				msgs = append(msgs, msg.Body)
			}

			res = append(res, &UserJSONResponse{
				Name:         user.Name,
				Email:        user.Email,
				SessionToken: user.SessionToken,
				Messages:     msgs,
			})
		}

		if bytes, err := json.Marshal(res); err != nil {
			renderError(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}
	}
}

// NewUserCreateHandler returns a handler that creates user.
func NewUserCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var regReq RegisterRequest
		if err := decoder.Decode(&regReq); err != nil {
			renderError(w, "Failed to parse request JSON into struct", http.StatusInternalServerError)
			return
		}

		if len(regReq.Email) == 0 || len(regReq.Password) == 0 || len(regReq.Name) == 0 {
			renderError(w, "Please provide name, email and password for registration", http.StatusBadRequest)
			return
		}

		hashBytes, hashErr := bcrypt.GenerateFromPassword([]byte(regReq.Password), 10)
		if hashErr != nil {
			renderError(w, hashErr.Error(), http.StatusInternalServerError)
			return
		}

		newUser := &model.User{
			Name:           regReq.Name,
			Email:          regReq.Email,
			PasswordDigest: hashBytes,
		}

		newUser.ResetSessionToken()

		if err := db.Create(&newUser).Error; err != nil {
			renderError(w, err.Error(), http.StatusInternalServerError)
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
			renderError(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}
	}
}
