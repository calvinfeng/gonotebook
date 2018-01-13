package main

import (
	"fmt"
	"go-academy/neural_net/layer"
	"gonum.org/v1/gonum/mat"
)

type NeuralNetwork struct {
	Layers     []layer.NetworkLayer
	BatchSize  int
	FeatureDim int
	HiddenDim  int
	OutputDim  int
}

func NewNeuralNetwork(batchSize, featureDim, hiddenDim, outputDim int) *NeuralNetwork {
	return &NeuralNetwork{
		Layers:     []layer.NetworkLayer{},
		BatchSize:  batchSize,
		FeatureDim: featureDim,
		HiddenDim:  hiddenDim,
		OutputDim:  outputDim,
	}
}

func (nn *NeuralNetwork) InitLayers(weightScale float64, numLayers int) {
	for i := 1; i <= numLayers; i += 1 {
		var nl layer.NetworkLayer

		switch i {
		case 1:
			nl = layer.NewAffineSigmoid(weightScale, nn.BatchSize, nn.FeatureDim, nn.HiddenDim)
		case numLayers:
			nl = layer.NewAffineSigmoid(weightScale, nn.BatchSize, nn.HiddenDim, nn.OutputDim)
		default:
			nl = layer.NewAffineSigmoid(weightScale, nn.BatchSize, nn.HiddenDim, nn.HiddenDim)
		}

		nn.Layers = append(nn.Layers, nl)
	}
}

func (nn *NeuralNetwork) TrainStep(X, Y *mat.Dense, learningRate float64) error {
	var forwardPropErr error
	Act := X
	for i := 0; i < len(nn.Layers); i += 1 {
		if Act, forwardPropErr = nn.Layers[i].ForwardProp(Act); forwardPropErr != nil {
			return forwardPropErr
		}
	}

	yRow, yCol := Y.Dims()
	outRow, outCol := Act.Dims()

	if yRow != outRow || yCol != outCol {
		return mat.ErrShape
	}

	UpstreamGrad := mat.NewDense(yRow, yCol, nil)
	UpstreamGrad.Sub(Act, Y)
	UpstreamGrad.Scale(2.0, UpstreamGrad)
	for i := len(nn.Layers) - 1; i >= 0; i = -1 {
		if bpResult, bpErr := nn.Layers[i].Update(learningRate, UpstreamGrad); bpErr != nil {
			return bpErr
		} else {
			UpstreamGrad = bpResult
		}
	}

	return nil
}

func (nn *NeuralNetwork) Loss(X, Y *mat.Dense) error {
	var forwardPropErr error
	Act := X
	for i := 0; i < len(nn.Layers); i += 1 {
		if Act, forwardPropErr = nn.Layers[i].ForwardProp(Act); forwardPropErr != nil {
			return forwardPropErr
		}
	}

	inRow, inCol := Y.Dims()
	outRow, outCol := Act.Dims()

	if inRow != outRow || inCol != outCol {
		return mat.ErrShape
	}

	loss := 0.0
	Diff := mat.NewDense(outRow, outCol, nil)
	Diff.Sub(Y, Act)

	for i := 0; i < outRow; i += 1 {
		for j := 0; j < outCol; j += 1 {
			loss += Diff.At(i, j) * Diff.At(i, j)
		}
	}

	fmt.Println(loss)
	return nil
}
