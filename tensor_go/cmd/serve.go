package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func serve(cmd *cobra.Command, args []string) error {
	fmt.Println("serving...")
	return nil
}
