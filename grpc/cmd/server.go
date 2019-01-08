package cmd

import (
	"fmt"
	"net"

	"go-academy/grpc/pb/planner"
	"go-academy/grpc/srv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const hostname = "localhost"
const httpPort = 8080
const rpcPort = 8081

// Server is a command for running server.
var Server = &cobra.Command{
	Use:   "server",
	Short: "Run Golang gRPC server",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	ch := make(chan error)

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", rpcPort))
		if err != nil {
			ch <- err
			return
		}

		logrus.Infof("launching gRPC server on port %d", rpcPort)
		gRPCServer := grpc.NewServer()

		planner.RegisterTodoServer(gRPCServer, srv.NewRPCTodoServer())

		ch <- gRPCServer.Serve(lis)
	}()

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", httpPort))
		if err != nil {
			ch <- err
			return
		}

		logrus.Infof("launching HTTP server on port %d", httpPort)
		httpServer := srv.NewHTTPTodoServer()
		httpServer.Register()

		ch <- httpServer.Serve(lis)
	}()

	return <-ch
}
