package main

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

func RandMat(row, col int) *mat.Dense {
	randFloats := []float64{}
	for i := 0; i < row*col; i++ {
		randFloats = append(randFloats, rand.Float64())
	}

	return mat.NewDense(row, col, randFloats)
}

func RandNormMat(row, col int, std, mean float64) *mat.Dense {
	randFloats := []float64{}
	for i := 0; i < row*col; i++ {
		randFloats = append(randFloats, rand.NormFloat64()*std+mean)
	}

	return mat.NewDense(row, col, randFloats)
}
