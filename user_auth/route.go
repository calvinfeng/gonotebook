package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"go-academy/user_auth/handler"
	"net/http"
)

// Gorilla mux library is a bit overkill for this example but it's good to introduce this powerful tool to you. Mux
// library offers URL pattern matching, query params patter matching, URL host matching and the list goes on.
// For example:
//
// r := mux.NewRouter()
// r.HandleFunc("/products/{key}", ProductHandler)
// r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
//
// Notice that key and category will become available as a variable through mux router pattern matching. If I were to
// send a request with the URL /products/1/, then key will hold the value 1.

func LoadRoutes(db *gorm.DB) http.Handler {
	// Defining middleware
	logMiddleware := NewServerLoggingMiddleware()

	// Instantiate our router object
	muxRouter := mux.NewRouter().StrictSlash(true)

	// Name-spacing the API
	api := muxRouter.PathPrefix("/api").Subrouter()
	api.Handle("/login", handler.NewSessionCreateHandler(db))
	api.Handle("/logout", handler.NewSessionDestroyHandler(db))

	// Serve public folder to clients
	muxRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	return handlers.CORS()(logMiddleware(muxRouter))
}
