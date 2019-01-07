package main

import (
	"go-academy/grpc/cmd"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:   "grpc",
		Short: "gRPC in Go",
	}

	root.AddCommand(cmd.Server, cmd.Client)

	if err := root.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
