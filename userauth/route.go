package main

import (
	"go-academy/userauth/handler"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Gorilla mux library is a bit overkill for this example but it's good to introduce this powerful
// tool to you. Mux library offers URL pattern matching, query params patter matching, URL host
// matching and the list goes on.
// For example:
//
// r := mux.NewRouter()
// r.HandleFunc("/products/{key}", ProductHandler)
// r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
//
// Notice that key and category will become available as a variable through mux router pattern
// matching. If I were to send a request with the URL /products/1/, then key will hold the value 1.

// HTTPMiddleware intercepts the a request before it reaches a handler.
type HTTPMiddleware func(http.Handler) http.Handler

func newServerLoggingMiddleware() HTTPMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logrus.Infof("%s %s %s %s", r.Proto, r.Method, r.URL, r.Host)
			next.ServeHTTP(w, r)
		})
	}
}

func loadRoutes(db *gorm.DB) http.Handler {
	// Defining middleware
	logMiddleware := newServerLoggingMiddleware()

	// Instantiate our router object
	muxRouter := mux.NewRouter().StrictSlash(true)

	// Name-spacing the API
	api := muxRouter.PathPrefix("/api").Subrouter()
	api.Handle("/users/login", handler.NewSessionCreateHandler(db)).Methods("POST")
	api.Handle("/users/logout", handler.NewSessionDestroyHandler(db)).Methods("DELETE")
	api.Handle("/users/authenticate", handler.NewTokenAuthenticateHandler(db)).Methods("GET")
	api.Handle("/users/register", handler.NewUserCreateHandler(db)).Methods("POST")
	api.Handle("/users", handler.NewUserListHandler(db)).Methods("GET")

	// Serve public folder to clients
	muxRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	return handlers.CORS()(logMiddleware(muxRouter))
}
