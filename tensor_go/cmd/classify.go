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

func classify(cmd *cobra.Command, args []string) error {
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

	tensor, err := tensorcv.GetTensorFromImagePath(imgPath, imgType, 3)
	if err != nil {
		return err
	}

	fmt.Println("Image tensor is loaded:", tensor.Shape())

	softmaxScore := tensorcv.RunResNetModel(tensor, ResNet)
	if softmaxScore == nil {
		return fmt.Errorf("unexpected problem occurred when resnet model is run, score is nil")
	}

	// Softmax score is of shape (N, 1000), since N is 1 here, so we will use the 0 indexed item.
	score := softmaxScore[0]
	top := make(map[int]bool)

	// Pick top 5, using the lazy way instead of writing a quick select...
	for i := 0; i < 5; i++ {
		classIdx, prob := tensorcv.ArgMax(score, top)
		fmt.Printf("Top predicted class is %s with %.2f probability\n", labels[classIdx], prob)
	}

	return nil
}
