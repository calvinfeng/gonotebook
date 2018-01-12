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

	batchSize, inputDim, hiddenDim, outputDim := 10, 4, 5, 3
	numLayers := 5
	layers := make(map[int]layer.NetworkLayer)
	weightScale := 1.0

	for i := 1; i <= numLayers; i += 1 {
		switch i {
		case 1:
			layers[i] = layer.NewAffineSigmoid(weightScale, batchSize, inputDim, hiddenDim)
		case numLayers:
			layers[i] = layer.NewAffineSigmoid(weightScale, batchSize, hiddenDim, outputDim)
		default:
			layers[i] = layer.NewAffineSigmoid(weightScale, batchSize, hiddenDim, hiddenDim)
		}
	}

	X := layer.RandNormMat(batchSize, inputDim, 1, 0)
	var LayerOut *mat.Dense
	var propErr error
	for i := 1; i <= numLayers; i += 1 {
		if i == 1 {
			LayerOut, propErr = layers[i].ForwardProp(X)
			if propErr != nil {
				break
			}
		} else {
			LayerOut, propErr = layers[i].ForwardProp(LayerOut)
			if propErr != nil {
				break
			}
		}
	}

	fmt.Println("Completed forward propagation")
	PrintMat(LayerOut)
}
