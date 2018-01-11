package layer

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

type Sigmoid struct {
	Input  *mat.Dense // Dense matrix of shape (N, D) where D is the feature dimension
	Output *mat.Dense // Dense matrix of shape (N, D)
}

func (s *Sigmoid) ForwardProp(x *mat.Dense) (*mat.Dense, error) {
	s.Input = x
	N, D := x.Dims()

	result := mat.NewDense(N, D, nil)
	for i := 0; i < N; i += 1 {
		for j := 0; j < D; j += 1 {
			result.Set(i, j, 1/(1+math.Exp(-1.0*x.At(i, j))))
		}
	}

	s.Output = result

	return result, nil
}

func (s *Sigmoid) BackwardProp(gradOut *mat.Dense) (*mat.Dense, error) {
	Nx, Dx := s.Output.Dims()
	N, D := gradOut.Dims()

	if Nx != N || Dx != D {
		return nil, mat.ErrShape
	}

	result := OnesMat(N, D)
	result.Sub(result, s.Output)
	result.MulElem(result, s.Output)
	result.MulElem(result, gradOut)

	return result, nil
}
