package cmd

import (
	"context"
	"fmt"
	"go-academy/grpc/pb/todo"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// Client is a command for running a gRPC client.
var Client = &cobra.Command{
	Use:   "client",
	Short: "Run gRPC client",
	RunE:  client,
}

func client(cmd *cobra.Command, args []string) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	bg := context.Background()
	conn, err := grpc.DialContext(bg, fmt.Sprintf("%s:%d", hostname, port), opts...)
	if err != nil {
		return err
	}

	defer conn.Close()

	todoCli := todo.NewTodoClient(conn)

	ctx, cancel := context.WithTimeout(bg, 5*time.Second)
	defer cancel()

	res, err := todoCli.Get(ctx, &todo.TodoRequest{Id: 1})
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
