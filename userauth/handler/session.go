package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/calvinfeng/go-academy/userauth/model"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

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

// NewSessionCreateHandler returns a handler that creates session for client. NOTE: Notice that I am
// not resetting the token during session creation, I will leave it to you as an exercise.
func NewSessionCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var loginReq LoginRequest
		if err := decoder.Decode(&loginReq); err != nil {
			renderError(w, "Failed to parse request JSON into struct", http.StatusInternalServerError)
			return
		}

		user, err := findUserByCredential(db, loginReq.Email, loginReq.Password)
		if err != nil {
			renderError(w, "Incorrect email/password combination", http.StatusUnauthorized)
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

		if bytes, err := json.Marshal(res); err != nil {
			renderError(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}
	}
}

func findUserByToken(db *gorm.DB, token string) (*model.User, error) {
	var user model.User
	if err := db.Where("session_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// NewSessionDestroyHandler returns a handler that destroys session for client.
func NewSessionDestroyHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Find current user using token from cookies
		cookie, _ := r.Cookie("session_token")
		if user, err := findUserByToken(db, cookie.Value); err == nil {
			user.ResetSessionToken()
			db.Save(user)

			res := &LogoutResponse{
				Name:        user.Name,
				Email:       user.Email,
				IsLoggedOut: true,
			}

			if bytes, err := json.Marshal(res); err != nil {
				renderError(w, err.Error(), http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(bytes)
			}
		} else {
			renderError(w, "User is not found", http.StatusBadRequest)
		}
	}
}

// NewTokenAuthenticateHandler returns a handler that authenticates user by returning a token as
// response.
func NewTokenAuthenticateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Find current user using token from cookies
		cookie, err := r.Cookie("session_token")
		if err != nil {
			renderError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if user, err := findUserByToken(db, cookie.Value); err == nil {
			res := UserJSONResponse{
				Name:         user.Name,
				Email:        user.Email,
				SessionToken: user.SessionToken,
			}

			if bytes, err := json.Marshal(res); err != nil {
				renderError(w, err.Error(), http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(bytes)
			}
		} else {
			renderError(w, err.Error(), http.StatusUnauthorized)
		}
	}
}
