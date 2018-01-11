package layer

import "gonum.org/v1/gonum/mat"

type NetworkLayer interface {
	// ForwardProp performs feed forward propagation of input and returns an output
	ForwardProp(*mat.Dense) (*mat.Dense, error)

	// Update performs gradient descent update on weight and return input gradient upon update success
	Update(*mat.Dense) (*mat.Dense, error)
}