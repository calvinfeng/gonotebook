package handler

import (
	"net/http"

	"github.com/calvinfeng/go-academy/userauth/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
)

// NewUserListHandler returns a handler that renders the list of users on the server.
func NewUserListHandler(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var users []*model.User

		if err := db.Preload("Messages").Find(&users).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}

		return ctx.JSON(http.StatusOK, users)
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

		user.Password = ""
		user.PasswordDigest = hashBytes
		user.ResetJWTToken()

		if err := db.Create(user).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusCreated, user)
	}
}

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

// NewCurrentUserRetrieveHandler returns a handler that fetches current user by token.
func NewCurrentUserRetrieveHandler(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, ctx.Get("current_user"))
	}
}
