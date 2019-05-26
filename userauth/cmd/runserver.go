package cmd

import (
	"io"
	"net/http"
	"os"

	"github.com/calvinfeng/go-academy/userauth/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"

	// Driver for Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// RunServerCmd is a command to run server from terminal.
var RunServerCmd = &cobra.Command{
	Use:   "runserver",
	Short: "run user authentication server",
	RunE:  runServer,
}

func runServer(cmd *cobra.Command, args []string) error {
	conn, err := gorm.Open("postgres", pgAddress)
	if err != nil {
		return err
	}

	srv := echo.New()

	srv.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "HTTP[${time_rfc3339}] ${method} ${path} status=${status} latency=${latency_human}\n",
		Output: io.MultiWriter(os.Stdout),
	}))

	srv.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	srv.File("/", "public/index.html")
	srv.POST("api/register/", handler.NewUserCreateHandler(conn))
	srv.POST("api/authenticate/", handler.NewUserAuthenticateHandler(conn))

	users := srv.Group("api/users")
	users.Use(handler.NewTokenAuthMiddleware(conn))
	users.GET("/", handler.NewUserListHandler(conn))

	if err := srv.Start(":8080"); err != nil {
		return err
	}

	return nil
}
