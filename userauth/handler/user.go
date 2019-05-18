package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/calvinfeng/go-academy/userauth/model"
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// NewUserListHandler returns a handler that renders the list of users on the server.
func NewUserListHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []model.User

		if err := db.Preload("Messages").Find(&users).Error; err != nil {
			logrus.WithField("src", "handler.user").Error(err)
			renderError(w, http.StatusInternalServerError, "database error")
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

		bytes, err := json.Marshal(res)
		if err != nil {
			logrus.WithField("src", "handler.user").Error(err)
			renderError(w, http.StatusInternalServerError, "failed to marshal response")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

// NewUserCreateHandler returns a handler that creates user.
func NewUserCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var regReq RegisterRequest
		if err := decoder.Decode(&regReq); err != nil {
			logrus.WithField("src", "handler.user").Error(err)
			renderError(w, http.StatusInternalServerError, "failed to unmarshal request data")
			return
		}

		if len(regReq.Email) == 0 || len(regReq.Password) == 0 || len(regReq.Name) == 0 {
			renderError(w, http.StatusBadRequest, "please provide name, email and password for registration")
			return
		}

		hashBytes, err := bcrypt.GenerateFromPassword([]byte(regReq.Password), 10)
		if err != nil {
			logrus.WithField("src", "handler.user").Error(err)
			renderError(w, http.StatusInternalServerError, "failed to generate hash bytes")
			return
		}

		newUser := &model.User{
			Name:           regReq.Name,
			Email:          regReq.Email,
			PasswordDigest: hashBytes,
		}

		newUser.ResetSessionToken()

		if err := db.Create(&newUser).Error; err != nil {
			logrus.WithField("src", "handler.user").Error(err)
			renderError(w, http.StatusInternalServerError, "database error")
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

		bytes, err := json.Marshal(res)
		if err != nil {
			logrus.WithField("src", "handler.user").Error(err)
			renderError(w, http.StatusInternalServerError, "failed to marshal response")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}
