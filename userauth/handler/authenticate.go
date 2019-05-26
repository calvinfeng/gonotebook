package handler

import (
	"net/http"

	"github.com/calvinfeng/go-academy/userauth/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
)

type (
	// Credential is a payload that captures user submitted credentials.
	Credential struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// TokenResponse is a payload that returns JWT token back to client.
	TokenResponse struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		JWTToken string `json:"jwt_token"`
	}
)

// NewUserAuthenticateHandler returns a handler that returns token.
func NewUserAuthenticateHandler(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		c := &Credential{}
		if err := ctx.Bind(c); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		user, err := findUserByCredentials(db, c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "wrong email or password")
		}

		return ctx.JSON(http.StatusOK, TokenResponse{
			Name:     user.Name,
			Email:    user.Email,
			JWTToken: user.JWTToken,
		})
	}
}

// NewUserCreateHandler returns a handler that creates user.
func NewUserCreateHandler(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := &model.User{}
		if err := ctx.Bind(user); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		if err := user.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
		}

		hashBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		user.PasswordDigest = hashBytes
		user.ResetJWTToken()

		if err := db.Create(user).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if err := ctx.JSON(http.StatusCreated, user); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return nil
	}
}
