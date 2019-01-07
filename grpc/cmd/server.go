package cmd

import (
	"fmt"
	"net"

	"go-academy/grpc/pb/todo"
	"go-academy/grpc/srv"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const hostname = "localhost"
const port = 8000

// Server is a command for running server.
var Server = &cobra.Command{
	Use:   "server",
	Short: "Run Golang gRPC server",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	gRPCServer := grpc.NewServer()
	todo.RegisterTodoServer(gRPCServer, srv.NewRPCTodoServer())

	return gRPCServer.Serve(lis)
}
