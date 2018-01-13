package layer

import (
	"gonum.org/v1/gonum/mat"
)

type NetworkLayer interface {
	// ForwardProp performs feed forward propagation of input and returns an output
	ForwardProp(*mat.Dense) (*mat.Dense, error)

	// Update performs gradient descent update on weight and return input gradient upon update success
	Update(float64, *mat.Dense) (*mat.Dense, error)

	// Retrieve weight
	Weight() *mat.Dense
}

type AffineSigmoid struct {
	AffLayer *Affine
	SigLayer *Sigmoid
}

// NewAffineSigmoid accepts three parameters, inDim (input dimension), and outDim (output dimension)
func NewAffineSigmoid(weightScale float64, inDim, outDim int) *AffineSigmoid {
	weightMat := RandNormMat(inDim, outDim, 1, 0)
	weightMat.Scale(weightScale, weightMat)

	return &AffineSigmoid{
		AffLayer: &Affine{
			Weight: weightMat,
			Bias:   mat.NewDense(1, outDim, nil),
		},
		SigLayer: &Sigmoid{},
	}
}

func (as *AffineSigmoid) ForwardProp(X *mat.Dense) (*mat.Dense, error) {
	if Score, affErr := as.AffLayer.ForwardProp(X); affErr == nil {
		if Act, sigErr := as.SigLayer.ForwardProp(Score); sigErr == nil {
			return Act, nil
		} else {
			return nil, sigErr
		}
	} else {
		return nil, affErr
	}
}

// GradAct - Gradient of Activations
func (as *AffineSigmoid) Update(learnRate float64, GradAct *mat.Dense) (*mat.Dense, error) {
	if GradScore, sigErr := as.SigLayer.BackwardProp(GradAct); sigErr == nil {
		if GradX, GradW, GradB, affErr := as.AffLayer.BackwardProp(GradScore); affErr == nil {
			// Perform update on weights
			GradW.Scale(learnRate, GradW)
			as.AffLayer.Weight.Sub(as.AffLayer.Weight, GradW)

			// Perform update on biases
			GradB.Scale(learnRate, GradB)
			as.AffLayer.Bias.Sub(as.AffLayer.Bias, GradB)

			return GradX, nil
		} else {
			return nil, affErr
		}
	} else {
		return nil, sigErr
	}
}

func (as *AffineSigmoid) Weight() *mat.Dense {
	N, H := as.AffLayer.Weight.Dims()

	m := mat.NewDense(N, H, nil)
	m.Copy(as.AffLayer.Weight)

	return m
}
