package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
)

// NewSessionCreateHandler returns a handler that creates session for client. NOTE: Notice that I am
// not resetting the token during session creation, I will leave it to you as an exercise.
func NewSessionCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var loginReq LoginRequest
		if err := decoder.Decode(&loginReq); err != nil {
			logrus.WithField("src", "handler.session").Error(err)
			renderError(w, http.StatusInternalServerError, "failed to unmarshal request data")
			return
		}

		user, err := findUserByCredential(db, loginReq.Email, loginReq.Password)
		if err != nil {
			logrus.WithField("src", "handler.session").Error(err)
			renderError(w, http.StatusUnauthorized, "incorrect email/password combination")
			return
		}

		// Token is set to expire in 2 days
		expiration := time.Now().Add(2 * 24 * time.Hour)
		cookie := http.Cookie{Name: "session_token", Value: user.SessionToken, Expires: expiration}
		http.SetCookie(w, &cookie)

		// Return the token anyways, just so we know for sure we got a legit token. Don't do this in production though...
		res := &UserJSONResponse{
			Name:         user.Name,
			Email:        user.Email,
			SessionToken: user.SessionToken,
		}

		bytes, err := json.Marshal(res)
		if err != nil {
			logrus.WithField("src", "handler.session").Error(err)
			renderError(w, http.StatusInternalServerError, "failed to marshal response")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

// NewSessionDestroyHandler returns a handler that destroys session for client.
func NewSessionDestroyHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Find current user using token from cookies
		cookie, err := r.Cookie("session_token")
		if err != nil {
			logrus.WithField("src", "handler.session").Error(err)
			renderError(w, http.StatusBadRequest, "session token is not present in cookie")
			return
		}

		user, err := findUserByToken(db, cookie.Value)
		if err != nil {
			logrus.WithField("src", "handler.session").Error(err)
			renderError(w, http.StatusBadRequest, "user is not found")
			return
		}

		user.ResetSessionToken()
		if err := db.Save(user).Error; err != nil {
			logrus.WithField("src", "handler.session").Error(err)
			renderError(w, http.StatusUnprocessableEntity, "database error")
			return
		}

		res := &LogoutResponse{
			Name:        user.Name,
			Email:       user.Email,
			IsLoggedOut: true,
		}

		bytes, err := json.Marshal(res)
		if err != nil {
			logrus.WithField("src", "handler.session").Error(err)
			renderError(w, http.StatusInternalServerError, "failed to marshal response")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

// NewTokenAuthenticateHandler returns a handler that authenticates user by returning a token as
// response.
func NewTokenAuthenticateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Find current user using token from cookies
		cookie, err := r.Cookie("session_token")
		if err != nil {
			logrus.WithField("src", "handler.session").Error(err)
			renderError(w, http.StatusUnauthorized, "session token is not found in cookie")
			return
		}

		user, err := findUserByToken(db, cookie.Value)
		if err != nil {
			logrus.WithField("src", "handler.session").Error(err)
			renderError(w, http.StatusUnauthorized, "cannot find user with provided token")
			return
		}

		res := UserJSONResponse{
			Name:         user.Name,
			Email:        user.Email,
			SessionToken: user.SessionToken,
		}

		bytes, err := json.Marshal(res)
		if err != nil {
			logrus.WithField("src", "handler.session").Error(err)
			renderError(w, http.StatusInternalServerError, "failed to marshal response")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}
