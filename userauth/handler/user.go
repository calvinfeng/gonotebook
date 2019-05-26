package handler

import (
	"net/http"

	"github.com/calvinfeng/go-academy/userauth/model"
	"github.com/labstack/echo/v4"

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
