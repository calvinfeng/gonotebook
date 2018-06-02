// Author(s): Calvin Feng

package cmd

import (
	"fmt"
	"go-academy/tensor_go/tensorcv"
	"os"

	"github.com/spf13/cobra"
)

// Model Path
const (
	ResNet = "./model/resnet"
)

var labels map[int]string

func init() {
	l, err := tensorcv.LoadLabels("./data/labels.json")
	if err != nil {
		panic(err)
	}

	labels = l
}

var rootCommand = &cobra.Command{
	Use:   "tensorgo",
	Short: "image recognition in Go using tensorflow and keras models",
}

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "start http server and listen for requests",
	RunE:  serve,
}

var classifyCommand = &cobra.Command{
	Use:   "classify",
	Short: "classify a specified png or jpeg image",
	RunE:  classify,
}

// Execute will activate the root command.
func Execute() {
	classifyCommand.Flags().String("img", "", "image path for the image you wish to classify")

	rootCommand.AddCommand(serveCommand, classifyCommand)
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
