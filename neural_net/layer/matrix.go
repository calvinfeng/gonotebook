package layer

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

func RandNormMat(row, col int, std, mean float64) *mat.Dense {
	rand.Seed(time.Now().UTC().Unix())

	randFloats := []float64{}
	for i := 0; i < row*col; i++ {
		randFloats = append(randFloats, rand.NormFloat64()*std+mean)
	}

	return mat.NewDense(row, col, randFloats)
}

func OnesMat(row, col int) *mat.Dense {
	ones := []float64{}
	for i := 0; i < row*col; i++ {
		ones = append(ones, 1)
	}

	return mat.NewDense(row, col, ones)
}
