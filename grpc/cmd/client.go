package cmd

import (
	"context"
	"fmt"
	"go-academy/grpc/cli"

	"github.com/sirupsen/logrus"
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

	t := cli.NewTodo(1)

	t.Networker = cli.NewHTTPNetworker("https://jsonplaceholder.typicode.com/todos")
	t.Load()

	logrus.Infof("fetched todo from HTTP, %v", t)

	t.Networker = cli.NewGRPCNetworker(conn)
	t.Load()

	t.Title = "Bye World"
	t.Save()

	t.Load()
	logrus.Infof("fetched todo from gRPC, %v", t)

	return nil
}
