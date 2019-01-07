package cmd

import (
	"net"

	"go-academy/grpc/pb/todo"
	"go-academy/grpc/srv"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// Server is a command for running server.
var Server = &cobra.Command{
	Use:   "server",
	Short: "Run Golang gRPC server",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		return err
	}

	gRPCServer := grpc.NewServer()
	todo.RegisterTodoServer(gRPCServer, srv.NewRPCTodoServer())

	return gRPCServer.Serve(lis)
}
