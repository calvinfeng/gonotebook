package handler

import (
	"net/http"

	"github.com/calvinfeng/go-academy/userauth/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// NewSentMessageListHandler returns a handler function that retrieves sent messages.
func NewSentMessageListHandler(db *gorm.DB) echo.HandlerFunc {
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
		if err := db.
			Preload("Sender").
			Preload("Receiver").
			Where("sender_id = ?", user.ID).Find(&messages).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		// Clear token for security
		for _, m := range messages {
			if m.Sender != nil {
				m.Sender.JWTToken = ""
			}

			if m.Receiver != nil {
				m.Receiver.JWTToken = ""
			}
		}

		return ctx.JSON(http.StatusOK, messages)
	}
}

// NewReceivedMessageListHandler returns a handler function that retrieves received messages.
func NewReceivedMessageListHandler(db *gorm.DB) echo.HandlerFunc {
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
		if err := db.
			Preload("Sender").
			Preload("Receiver").
			Where("receiver_id = ?", user.ID).Find(&messages).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		// Clear token for security
		for _, m := range messages {
			if m.Sender != nil {
				m.Sender.JWTToken = ""
			}

			if m.Receiver != nil {
				m.Receiver.JWTToken = ""
			}
		}

		return ctx.JSON(http.StatusOK, messages)
	}
}

// NewMessageCreateHandler returns a handler function that creates message.
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

		message.SenderID = user.ID
		if err := message.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
		}

		if err := db.Create(message).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusCreated, message)
	}
}
