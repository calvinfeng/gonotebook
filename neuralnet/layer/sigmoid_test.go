package layer

import (
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
	"math"
	"testing"
)

//SigForwardProp := func(input *mat.Dense) (*mat.Dense, error) {
//	sig := layer.Sigmoid{}
//
//	return sig.ForwardProp(input)
//}

func TestSigmoid(t *testing.T) {
	// X is an input to the sigmoid layer, with shape (3, 3)
	X := RandNormMat(3, 3, 1, 0)

	// grad is the upstream gradient, it is dL/d_output
	grad := OnesMat(3, 3)

	sig := Sigmoid{}

	t.Run("TestForwardProp", func(t *testing.T) {
		if output, err := sig.ForwardProp(X); err != nil {
			t.Error("Failed to perform forward propagation", err)
		} else {
			row, col := output.Dims()

			expected := mat.NewDense(3, 3, nil)
			for i := 0; i < row; i += 1 {
				for j := 0; j < col; j += 1 {
					expected.Set(i, j, 1/(1+math.Exp(-1.0*X.At(i, j))))
				}
			}

			assert.Equal(t, 3, row, "Row count should equal to 3")
			assert.Equal(t, 3, col, "Col count should equal to 3")

			for i := 0; i < row; i += 1 {
				for j := 0; j < col; j += 1 {
					assert.Equal(t, expected.At(i, j), output.At(i, j), "Output should equal to expected")
				}
			}
		}
	})

	t.Run("TestBackwardProp", func(t *testing.T) {
		if gradX, err := sig.BackwardProp(grad); err != nil {
			t.Error("Failed to perform back propagation", err)
		} else {
			// Compute numerical gradient with respect to X
			numGradX, numErr := EvalNumericalGrad(sig.ForwardProp, X, grad, 1e-5)
			if numErr != nil {
				t.Error("Failed to evaluate numerical gradient on X", numErr)
			}

			relError := mat.NewDense(3, 3, nil)
			relError.Sub(gradX, numGradX)

			row, col := relError.Dims()
			for i := 0; i < row; i += 1 {
				for j := 0; j < col; j += 1 {
					assert.Equal(t, float64(int(relError.At(i, j)*1e8)/1e8), 0.0, "Error should be zero")
				}
			}
		}
	})
}
