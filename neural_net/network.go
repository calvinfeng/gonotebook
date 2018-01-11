package main

import (
	"go-academy/neural_net/layer"
)

type NeuralNetwork struct {
	Layers     []layer.NetworkLayer
	BatchSize  int
	FeatureDim int
	HiddenDim  int
}

func NewNeuralNetwork(batchSize, featureDim, hiddenDim int) *NeuralNetwork {
	return &NeuralNetwork{
		Layers:     []layer.NetworkLayer{},
		BatchSize:  batchSize,
		FeatureDim: featureDim,
		HiddenDim:  hiddenDim,
	}
}

func InitLayers(numLayers int) {

}
