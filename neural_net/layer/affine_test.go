package layer

import (
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestAffine(t *testing.T) {
	// X is an input to the affine layer, with shape (3, 3)
	X := RandNormMat(3, 3, 1, 0)

	// W is a weight matrix assigned to the affine layer
	W := RandNormMat(3, 3, 1, 0)

	// B is a bias matrix
	B := RandNormMat(1, 3, 1e-2, 0)

	// grad is the upstream gradient, it is dL/d_output
	grad := OnesMat(3, 3)

	aff := Affine{
		Weight: W,
		Bias:   B,
	}

	t.Run("TestForwardProp", func(t *testing.T) {
		if output, err := aff.ForwardProp(X); err != nil {
			t.Error("Failed to perform forward propagation", err)
		} else {
			expected := mat.NewDense(3, 3, nil)
			expected.Mul(X, W)
			expected.Add(expected, MatBroadcast(B, 3))

			row, col := output.Dims()
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
		if gradX, gradW, gradB, err := aff.BackwardProp(grad); err != nil {
			t.Error("Failed to perform back propagation", err)
		} else {
			var numGradX, numGradW, numGradB, relError *mat.Dense
			var row, col int
			var numErr error

			// Compute numerical gradient with respect to X
			numGradX, numErr = EvalNumericalGrad(aff.ForwardProp, X, grad, 1e-5)
			if numErr != nil {
				t.Error("Failed to evaluate numerical gradient on X", numErr)
			}

			relError = mat.NewDense(3, 3, nil)
			relError.Sub(gradX, numGradX)

			row, col = relError.Dims()
			for i := 0; i < row; i += 1 {
				for j := 0; j < col; j += 1 {
					assert.Equal(t, float64(int(relError.At(i, j)*1e8)/1e8), 0.0, "Error should be zero")
				}
			}

			// Compute numerical gradient with respect to W
			forwardPropW := func(weight *mat.Dense) (*mat.Dense, error) {
				aff.Weight = weight

				// Don't forget to reset
				defer func() {
					aff.Weight = W // Reset it back to original value
				}()

				if output, err := aff.ForwardProp(X); err == nil {
					return output, nil
				} else {
					return nil, err
				}
			}

			numGradW, numErr = EvalNumericalGrad(forwardPropW, W, grad, 1e-5)
			if numErr != nil {
				t.Error("Failed to evaluate numerical gradient on W", numErr)
			}

			relError = mat.NewDense(3, 3, nil)
			relError.Sub(gradW, numGradW)

			row, col = relError.Dims()
			for i := 0; i < row; i += 1 {
				for j := 0; j < col; j += 1 {
					assert.Equal(t, float64(int(relError.At(i, j)*1e8)/1e8), 0.0, "Error should be zero")
				}
			}

			// Compute numerical gradient with respect to B
			forwardPropB := func(bias *mat.Dense) (*mat.Dense, error) {
				aff.Bias = bias

				// Don't forget to reset
				defer func() {
					aff.Bias = B
				}()

				if output, err := aff.ForwardProp(X); err == nil {
					return output, nil
				} else {
					return nil, err
				}
			}

			numGradB, numErr = EvalNumericalGradForBias(forwardPropB, B, grad, 1e-5)
			if numErr != nil {
				t.Error("Failed to evaluate numerical gradient on B", numErr)
			}

			relError = mat.NewDense(1, 3, nil)
			relError.Sub(gradB, numGradB)

			row, col = relError.Dims()
			if row != 1 {
				t.Error("Bias row dimension is not one")
			}

			for j := 0; j < col; j += 1 {
				assert.Equal(t, float64(int(relError.At(0, j)*1e8)/1e8), 0.0, "Error should be zero")
			}
		}
	})
}
