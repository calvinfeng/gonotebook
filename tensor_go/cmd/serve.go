package cmd

import (
	"go-academy/tensor_go/tensorcv"

	"github.com/spf13/cobra"
)

func serve(cmd *cobra.Command, args []string) error {
	tensorcv.HelloWorldFromTF()
	return nil
}
