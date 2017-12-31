package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func LoadRoutes(db *gorm.DB) http.Handler {
	// Defining middleware
	logMiddleware := NewServerLoggingMiddleware()

	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	// Name-spacing the API
	// api := muxRouter.PathPrefix("/api").Subrouter()

	return handlers.CORS()(logMiddleware(muxRouter))
}
