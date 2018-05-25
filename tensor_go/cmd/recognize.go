package cmd

import (
	"errors"
	"fmt"
	"go-academy/tensor_go/tensorcv"
	"strings"

	"github.com/spf13/cobra"
)

// Image file extensions
const (
	PNG  = "png"
	JPG  = "jpg"
	JPEG = "jpeg"
)

func recognize(cmd *cobra.Command, args []string) error {
	imgPath, err := cmd.Flags().GetString("img")
	if err != nil {
		return err
	}

	if imgPath == "" {
		return errors.New("please specify a valid image file location")
	}

	splitPath := strings.Split(imgPath, ".")

	var imgType string
	switch splitPath[len(splitPath)-1] {
	case PNG:
		imgType = "png"
	case JPEG:
		imgType = "jpeg"
	case JPG:
		imgType = "jpeg"
	default:
		return fmt.Errorf("%s is not a valid image", imgPath)
	}

	tensor, err := tensorcv.GetTensorFromImagePath(imgPath, imgType, 4)
	fmt.Println("Image tensor is loaded:", tensor.Shape())

	softmaxScore := tensorcv.RunResNetModel(tensor)
	if softmaxScore == nil {
		return fmt.Errorf("unexpected problem occurred when resnet model is run, score is nil")
	}

	return nil
}
