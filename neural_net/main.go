package main

import (
	"fmt"
	"go-academy/neural_net/layer"
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

func PrintMat(M *mat.Dense) {
	row, col := M.Dims()
	line := ""
	for i := 0; i < row; i += 1 {
		for j := 0; j < col; j += 1 {
			line += fmt.Sprintf(" %v ", M.At(i, j))
		}
		line += "\n"
	}

	fmt.Println(line)
}

func main() {
	// If seed is not set, we will expect to get the result random numbers in runtime every time we run the program. This
	// particular seed will give the program pseudo-randomness.
	rand.Seed(time.Now().UTC().Unix())

	batchSize, inputDim, hiddenDim, outputDim, numLayers, weightScale := 10, 4, 10, 3, 5, 1e-2

	X := layer.RandNormMat(batchSize, inputDim, 1, 0)
	Y := layer.RandNormMat(batchSize, outputDim, 1, 0)

	network := NewNeuralNetwork(batchSize, inputDim, hiddenDim, outputDim)
	network.InitLayers(weightScale, numLayers)

	if err := network.Loss(X, Y); err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 10000; i += 1 {
		if err := network.TrainStep(X, Y, 1e-2); err != nil {
			fmt.Println(err)
			break
		}
	}

	if err := network.Loss(X, Y); err != nil {
		fmt.Println(err)
	}
}
