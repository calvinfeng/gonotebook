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
	conn, err := grpc.DialContext(bg, fmt.Sprintf("%s:%d", hostname, rpcPort), opts...)
	if err != nil {
		return err
	}

	defer conn.Close()

	httpNet := cli.NewHTTPNetworker(fmt.Sprintf("http://%s:%d/todos", hostname, httpPort))
	grpcNet := cli.NewGRPCNetworker(conn)

	t := cli.NewTodo(1, httpNet)
	t.Title = "Hello World"
	t.Completed = true
	t.UserID = 1

	logrus.Info("save todo via HTTP")
	if err := t.Save(); err != nil {
		logrus.Error(err)
	}

	t = cli.NewTodo(1, httpNet)
	t.Load()
	logrus.Infof("fetched todo via HTTP, %v", t)

	t = cli.NewTodo(1, grpcNet)
	t.Load()
	logrus.Infof("fetched todo via gRPC, %v", t)

	return nil
}
