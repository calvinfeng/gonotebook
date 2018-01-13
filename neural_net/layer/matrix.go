package layer

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

func RandMat(row, col int) *mat.Dense {
	rand.Seed(time.Now().UTC().Unix())

	randFloats := []float64{}
	for i := 0; i < row*col; i++ {
		randFloats = append(randFloats, rand.Float64())
	}

	return mat.NewDense(row, col, randFloats)
}

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

// MatBroadcast is analogous to the NumPy array broadcasting. It takes a (1, H) matrix and return a (N, H) matrix.
func MatBroadcast(m *mat.Dense, N int) *mat.Dense {
	_, H := m.Dims()
	stack := mat.NewDense(0, H, nil)

	for i := 1; i <= N; i += 1 {
		temp := mat.NewDense(i, H, nil)
		temp.Stack(stack, m)
		stack = temp
	}

	return stack
}

// SumAlongColumn accepts a matrix and perform summing along each column of the matrix. If the input is of shape (N, D),
// the return value will be a slice of length D.
func SumAlongColumn(m *mat.Dense) []float64 {
	Row, Col := m.Dims()
	result := []float64{}
	for j := 0; j < Col; j += 1 {
		colSum := 0.0
		for i := 0; i < Row; i += 1 {
			colSum += m.At(i, j)
		}
		result = append(result, colSum)
	}

	return result
}
