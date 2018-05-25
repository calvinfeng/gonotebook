// Author(s): Calvin Feng

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "tensorgo",
	Short: "image recognition in Go using tensorflow and keras models",
}

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "start http server and listen for requests",
	RunE:  serve,
}

var recognizeCommand = &cobra.Command{
	Use:   "classify",
	Short: "classify a specified png or jpeg image",
	RunE:  recognize,
}

// Execute will activate the root command.
func Execute() {
	recognizeCommand.Flags().String("img", "", "image path for the image you wish to classify")

	rootCommand.AddCommand(serveCommand, recognizeCommand)
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
