package handler

import (
	"encoding/json"
	"net/http"
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

func renderError(w http.ResponseWriter, errMsg string, code int) {
	res := &ErrorResponse{errMsg}
	bytes, _ := json.Marshal(res)
	w.WriteHeader(code)
	w.Write(bytes)
}
