package tensorcv

import (
	"fmt"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

// RunResNetModel ...
func RunResNetModel(imgTensor *tf.Tensor) [][]float32 {
	model, err := tf.LoadSavedModel("./tf_resnet_py/model", []string{"serve"}, nil)
	if err != nil {
		fmt.Printf("Error loading saved model: %s\n", err.Error())
		return nil
	}

	defer model.Session.Close()

	feeds := map[tf.Output]*tf.Tensor{
		model.Graph.Operation("input_1").Output(0): imgTensor,
	}

	fmt.Println("Number of available outputs:", model.Graph.Operation("fc1000/Softmax").NumOutputs())
	fetches := []tf.Output{model.Graph.Operation("fc1000/Softmax").Output(0)}

	result, runErr := model.Session.Run(feeds, fetches, nil)
	if runErr != nil {
		fmt.Printf("Error running the session with input, err: %s\n", runErr.Error())
		return nil
	}

	return result[0].Value().([][]float32)
}

// RunMnistModel ...
func RunMnistModel(imgTensor *tf.Tensor) {
	model, err := tf.LoadSavedModel("./tf_mnist_py/model", []string{"serve"}, nil)
	if err != nil {
		fmt.Printf("Error loading saved model: %s\n", err.Error())
		return
	}

	defer model.Session.Close()

	feeds := map[tf.Output]*tf.Tensor{
		model.Graph.Operation("image_input").Output(0): imgTensor,
	}

	fetches := []tf.Output{model.Graph.Operation("infer").Output(0)}

	result, runErr := model.Session.Run(feeds, fetches, nil)
	if runErr != nil {
		fmt.Printf("Error running the session with input, err: %s\n", runErr.Error())
		return
	}

	fmt.Printf("Most likely number in input is %v\n", result[0].Value())
}
