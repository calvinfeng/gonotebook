package main

import (
	"fmt"
	"go-academy/neural_net/layer"
	"gonum.org/v1/gonum/mat"
)

type NeuralNetwork struct {
	Layers    []layer.NetworkLayer
	InputDim  int
	HiddenDim int
	OutputDim int
}

func NewNeuralNetwork(inputDim, hiddenDim, outputDim int) *NeuralNetwork {
	return &NeuralNetwork{
		Layers:    []layer.NetworkLayer{},
		InputDim:  inputDim,
		HiddenDim: hiddenDim,
		OutputDim: outputDim,
	}
}

func (nn *NeuralNetwork) InitLayers(weightScale float64, numLayers int) {
	for i := 1; i <= numLayers; i += 1 {
		var nl layer.NetworkLayer

		switch i {
		case 1:
			nl = layer.NewAffineSigmoid(weightScale, nn.InputDim, nn.HiddenDim)
		case numLayers:
			nl = layer.NewAffineSigmoid(weightScale, nn.HiddenDim, nn.OutputDim)
		default:
			nl = layer.NewAffineSigmoid(weightScale, nn.HiddenDim, nn.HiddenDim)
		}

		nn.Layers = append(nn.Layers, nl)
	}
}

func (nn *NeuralNetwork) PrintWeights() {
	for i := 0; i < len(nn.Layers); i += 1 {
		PrintMat(nn.Layers[i].Weight())
	}
}

func (nn *NeuralNetwork) TrainStep(X, Y *mat.Dense, learningRate float64) error {
	Score, loss, err := nn.Loss(X, Y)
	if err != nil {
		return err
	}

	yRow, yCol := Y.Dims()
	outRow, outCol := Score.Dims()

	if yRow != outRow || yCol != outCol {
		return mat.ErrShape
	}

	fmt.Printf("Loss during training step: %.5f\n", loss)

	UpstreamGrad := mat.NewDense(yRow, yCol, nil)
	UpstreamGrad.Sub(Score, Y)
	UpstreamGrad.Scale(2.0, UpstreamGrad)
	for i := len(nn.Layers) - 1; i >= 0; i -= 1 {
		if bpResult, bpErr := nn.Layers[i].Update(learningRate, UpstreamGrad); bpErr != nil {
			return bpErr
		} else {
			UpstreamGrad = bpResult
		}
	}

	return nil
}

func (nn *NeuralNetwork) Loss(X, Y *mat.Dense) (*mat.Dense, float64, error) {
	var forwardPropErr error
	Act := X
	for i := 0; i < len(nn.Layers); i += 1 {
		if Act, forwardPropErr = nn.Layers[i].ForwardProp(Act); forwardPropErr != nil {
			return nil, 0, forwardPropErr
		}
	}
	Score := Act

	inRow, inCol := Y.Dims()
	outRow, outCol := Act.Dims()

	if inRow != outRow || inCol != outCol {
		return nil, 0, mat.ErrShape
	}

	loss := 0.0
	Diff := mat.NewDense(outRow, outCol, nil)
	Diff.Sub(Y, Score)

	for i := 0; i < outRow; i += 1 {
		for j := 0; j < outCol; j += 1 {
			loss += Diff.At(i, j) * Diff.At(i, j)
		}
	}

	return Score, loss, nil
}
