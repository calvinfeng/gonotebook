package layer

import (
	"gonum.org/v1/gonum/mat"
)

type Affine struct {
	Input  *mat.Dense // Dense matrix of shape (N, D) where D is the feature dimension
	Weight *mat.Dense // Dense matrix of shape (D, H) where H is the hidden dimension
	Bias   *mat.Dense // Dense matrix of shape (N, H)
	Output *mat.Dense // Dense matrix of shape (N, H)
}

// Expected input is of shape (N, D) where N is the number of batch example
func (a *Affine) ForwardProp(x *mat.Dense) (*mat.Dense, error) {
	// First check dimension
	xRow, xCol := x.Dims()
	wRow, wCol := a.Weight.Dims()
	bRow, bCol := a.Bias.Dims()

	if xCol != wRow || xRow != bRow || wCol != bCol {
		return nil, mat.ErrShape
	}

	// Cache the input, cause we need it to compute back propagation
	a.Input = mat.NewDense(xRow, xCol, nil)
	a.Input.Copy(x)

	// Initialize a nil matrix
	result := mat.NewDense(xRow, wCol, nil)

	// Perform operations
	result.Mul(x, a.Weight)
	result.Add(result, a.Bias)

	return result, nil
}

// BackwardProp takes an upstream gradient, a.k.a. gradient of output and perform gradient calculation on layer's input,
// weight and bias matrices. The expected input to this function should be of shape, (N, H) which is the same shape as
// the output from ForwardProp.
func (a *Affine) BackwardProp(gradOut *mat.Dense) (*mat.Dense, *mat.Dense, *mat.Dense, error) {
	if a.Input == nil {
		return nil, nil, nil, mat.Error{}
	}

	gradOutRow, gradOutCol := gradOut.Dims()
	xRow, xCol := a.Input.Dims()
	wRow, wCol := a.Weight.Dims()

	if gradOutRow != xRow || gradOutCol != wCol {
		return nil, nil, nil, mat.ErrShape
	}

	gradW := mat.NewDense(xCol, gradOutCol, nil)
	gradW.Mul(a.Input.T(), gradOut)

	gradX := mat.NewDense(gradOutRow, wRow, nil)
	gradX.Mul(gradOut, a.Weight.T())

	// Sum along column on upstream gradient
	sumSlice := SumAlongColumn(gradOut)
	biasData := []float64{}
	for i := 0; i < gradOutRow; i += 1 {
		biasData = append(biasData, sumSlice...)
	}

	gradB := mat.NewDense(gradOutRow, gradOutCol, biasData)

	return gradX, gradW, gradB, nil
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
