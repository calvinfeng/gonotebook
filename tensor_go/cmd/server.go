package cmd

import (
	"fmt"
	"go-academy/tensor_go/tensorcv"
	"net/http"

	"github.com/spf13/cobra"
)

func server(cmd *cobra.Command, args []string) error {
	port := ":3000"

	if len(args) > 0 {
		port = fmt.Sprintf(":%s", args[0])
	}

	server := &http.Server{
		Addr:    port,
		Handler: tensorcv.LoadRoutes(labels, ResNet),
	}

	fmt.Printf("server is listening and serving on %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
