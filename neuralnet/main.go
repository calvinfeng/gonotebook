package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

func Argmax(m *mat.Dense) []int {
	N, C := m.Dims()

	result := []int{}
	for i := 0; i < N; i += 1 {
		maxScore := m.At(i, 0)
		maxIdx := 0
		for j := 0; j < C; j += 1 {
			if m.At(i, j) > maxScore {
				maxScore = m.At(i, j)
				maxIdx = j
			}
		}

		result = append(result, maxIdx)
	}

	return result
}

func Accuracy(predScore *mat.Dense, expectScore *mat.Dense) (float64, error) {
	predRow, predCol := predScore.Dims()
	exptRow, exptCol := expectScore.Dims()

	if predRow != exptRow || predCol != exptCol {
		return 0, mat.ErrShape
	}

	numCorrect := 0.0
	pred := Argmax(predScore)
	expt := Argmax(expectScore)

	for i := 0; i < len(pred); i += 1 {
		if pred[i] == expt[i] {
			numCorrect += 1.0
		}
	}

	return numCorrect / float64(len(pred)), nil
}

func PrintMat(M *mat.Dense) {
	row, col := M.Dims()
	line := ""
	for i := 0; i < row; i += 1 {
		for j := 0; j < col; j += 1 {
			line += fmt.Sprintf(" %.3f ", M.At(i, j))
		}
		line += "\n"
	}

	fmt.Println(line)
}

func main() {
	// If seed is not set, we will expect to get the result random numbers in runtime every time we run the program. This
	// particular seed will give the program pseudo-randomness.
	rand.Seed(time.Now().UTC().Unix())

	hiddenDim, numLayers, weightScale := 5, 2, 1e-2

	Xtr, Ytr := LoadIrisData("data/iris_train.csv")
	Xval, Yval := LoadIrisData("data/iris_val.csv")

	_, inputDim := Xtr.Dims()
	_, outputDim := Ytr.Dims()

	network := NewNeuralNetwork(inputDim, hiddenDim, outputDim)
	network.InitLayers(weightScale, numLayers)
	network.PrintWeights()

	for i := 0; i < 1000; i += 1 {
		if err := network.TrainStep(Xtr, Ytr, 1e-2); err != nil {
			fmt.Println(err)
			break
		}
	}

	network.PrintWeights()

	if Score, loss, err := network.Loss(Xval, Yval); err == nil {
		fmt.Printf("Completed training, validation loss: %.5f\n", loss)

		if acc, err := Accuracy(Score, Yval); err == nil {
			fmt.Printf("Accuracy is %.5f\n", acc)
		}
	}
}
