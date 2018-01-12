package layer

import "gonum.org/v1/gonum/mat"

// Any function that accepts a Dense matrix and returns a Dense matrix will satisfy the definition of a ForwardProp
type ForwardProp func(*mat.Dense) (*mat.Dense, error)

// EvalNumericalGradient accepts four arguments, the h represents a small change in input and it can be either positive
// or negative.
func EvalNumericalGrad(f ForwardProp, input *mat.Dense, upstreamGrad *mat.Dense, h float64) (*mat.Dense, error) {
	inRow, inCol := input.Dims()
	outRow, outCol := upstreamGrad.Dims()

	if inRow != outRow || inCol != outCol {
		return nil, mat.ErrShape
	}

	numGradient := mat.NewDense(inRow, inCol, nil)
	for i := 0; i < inRow; i += 1 {
		for j := 0; j < inCol; j += 1 {
			var fxph, fxmh *mat.Dense // f(x + h) and f(x - h)

			oldVal := input.At(i, j)

			input.Set(i, j, oldVal+h)
			if output, err := f(input); err == nil {
				fxph = output
			} else {
				return nil, err
			}

			input.Set(i, j, oldVal-h)
			if output, err := f(input); err == nil {
				fxmh = output
			} else {
				return nil, err
			}

			// Reset the input back to its original value
			input.Set(i, j, oldVal)

			diff := mat.NewDense(outRow, outCol, nil)
			diff.Sub(fxph, fxmh)
			diff.MulElem(diff, upstreamGrad)

			numGradient.Set(i, j, mat.Sum(diff)/(2.0*h))
		}
	}

	return numGradient, nil
}

func EvalNumericalGradForBias(f ForwardProp, input *mat.Dense, upstreamGrad *mat.Dense, h float64) (*mat.Dense, error) {
	inRow, inCol := input.Dims()
	outRow, outCol := upstreamGrad.Dims()

	if inRow != outRow || inCol != outCol {
		return nil, mat.ErrShape
	}

	numGradient := mat.NewDense(inRow, inCol, nil)
	for j := 0; j < inCol; j += 1 {
		// Sum all the numerical slope values along j column, i.e. iterating from top to bottom i = 0, 1, ..., inRow - 1
		slopeSum := 0.0
		for i := 0; i < inRow; i += 1 {
			var fxph, fxmh *mat.Dense // f(x + h) and f(x - h)

			oldVal := input.At(i, j)

			input.Set(i, j, oldVal+h)
			if output, err := f(input); err == nil {
				fxph = output
			} else {
				return nil, err
			}

			input.Set(i, j, oldVal-h)
			if output, err := f(input); err == nil {
				fxmh = output
			} else {
				return nil, err
			}

			// Reset the input back to its original value
			input.Set(i, j, oldVal)

			diff := mat.NewDense(outRow, outCol, nil)
			diff.Sub(fxph, fxmh)
			diff.MulElem(diff, upstreamGrad)

			slopeSum += mat.Sum(diff) / (2.0 * h)
		}

		// Set all value in the column of the numerical gradient to this slope sum
		for i := 0; i < inRow; i += 1 {
			numGradient.Set(i, j, slopeSum)
		}
	}

	return numGradient, nil
}
