package handler

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type SessionJSONResponse struct {
	Email        string `json:"email"`
	SessionToken string `json:"session_token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NOTE: Notice that I am not resetting the token during session creation, I will leave it to you as an exercise.
func NewSessionCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var loginReq LoginRequest
		if err := decoder.Decode(&loginReq); err != nil {
			RenderError(w, "Failed to parse request JSON into struct", http.StatusInternalServerError)
			return
		}

		user, err := FindUserByCredential(db, loginReq.Email, loginReq.Password)
		if err != nil {
			RenderError(w, "Wrong email/password combination", http.StatusUnauthorized)
			return
		}

		// Token is set to expire in 2 days
		expiration := time.Now().Add(2 * 24 * time.Hour)
		cookie := http.Cookie{Name: "session_token", Value: user.SessionToken, Expires: expiration}
		http.SetCookie(w, &cookie)

		// Return the token anyways, just so we know for sure we got a legit token. Don't do this in production though...
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

func NewSessionDestroyHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Find current user using token from cookies
		cookie, _ := r.Cookie("session_token")
		if currentUser, err := FindUserByToken(db, cookie.Value); err == nil {
			currentUser.ResetSessionToken()
			db.Save(currentUser)

			// Set the cookie to nil value once session is destroyed
			cookie := http.Cookie{Name: "session_token", Value: ""}
			http.SetCookie(w, &cookie)
		}
	}
}
