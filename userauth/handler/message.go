package handler

import (
	"net/http"

	"github.com/calvinfeng/go-academy/userauth/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func NewMessageListByCurrentUser(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Get("current_user")
		if val == nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to find current user")
		}

		user, ok := val.(*model.User)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to find current user")
		}

		var messages []*model.Message
		if err := db.Where("user_id = ?", user.ID).Find(&messages).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return ctx.JSON(http.StatusOK, messages)
	}
}

func NewMessageCreateHandler(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Get("current_user")
		if val == nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to find current user")
		}

		user, ok := val.(*model.User)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to find current user")
		}

		message := &model.Message{}
		if err := ctx.Bind(message); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		message.UserID = user.ID
		if err := message.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
		}

		if err := db.Create(message).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusCreated, message)
	}
}
