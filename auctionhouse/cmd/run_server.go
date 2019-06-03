package cmd

import (
	"io"
	"net"
	"net/http"
	"os"

	"github.com/calvinfeng/go-academy/auctionhouse/auction"
	"github.com/calvinfeng/go-academy/auctionhouse/protobuf"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	// Driver for Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// RunServerCmd is a command to run both gRPC and HTTP servers.
var RunServerCmd = &cobra.Command{
	Use:   "runserver",
	Short: "run auction HTTP/gRPC server",
	RunE:  runServer,
}

func runServer(cmd *cobra.Command, args []string) error {
	_, err := gorm.Open("postgres", pgAddr)
	if err != nil {
		return err
	}

	errChan := make(chan error)
	go func() {
		srv := echo.New()

		srv.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "HTTP[${time_rfc3339}] ${method} ${path} status=${status} latency=${latency_human}\n",
			Output: io.MultiWriter(os.Stdout),
		}))

		srv.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))

		errChan <- srv.Start(":8080")
	}()

	go func() {
		gRPCLis, err := net.Listen("tcp", ":8081")
		if err != nil {
			errChan <- err
		}

		gRPCSrv := grpc.NewServer()
		protobuf.RegisterAuctionServer(gRPCSrv, &auction.ServiceServer{})

		errChan <- gRPCSrv.Serve(gRPCLis)
	}()

	return <-errChan
}
