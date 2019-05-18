package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/calvinfeng/go-academy/userauth/handler"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	// Driver for Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const addr = ":3000"

func connectDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s",
			user, password, host, port, database, sslMode),
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}

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

func runserver(cmd *cobra.Command, args []string) error {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	db, err := connectDB()
	if err != nil {
		return err
	}

	defer db.Close()

	server := &http.Server{
		Handler:      loadRoutes(db),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Infof("HTTP server is listening and serving on port %v", addr)
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
