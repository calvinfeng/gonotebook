package tensorcv

import (
	"bytes"
	"io"
	"os"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

// Constants for scaling images.
const (
	ImageHeight = 227
	ImageWidth  = 227
	Mean        = float32(0)
	Scale       = float32(1)
)

// GetTensorFromImagePath creates a tensor struct by taking an image path.
func GetTensorFromImagePath(imgPath, imgFormat string, numChan int64) (*tf.Tensor, error) {
	var err error

	imgFile, _ := os.Open(imgPath)
	defer imgFile.Close()

	var imgBuffer bytes.Buffer
	io.Copy(&imgBuffer, imgFile)

	var imgTensor *tf.Tensor
	imgTensor, err = tf.NewTensor(imgBuffer.String())
	if err != nil {
		return nil, err
	}

	var graph *tf.Graph
	var input, output tf.Output
	graph, input, output, err = createImageTransformGraph(imgFormat, numChan)
	if err != nil {
		return nil, err
	}

	var sess *tf.Session
	sess, err = tf.NewSession(graph, nil)
	if err != nil {
		return nil, err
	}

	defer sess.Close()

	feeds := map[tf.Output]*tf.Tensor{
		input: imgTensor,
	}

	fetches := []tf.Output{output}
	if normalized, err := sess.Run(feeds, fetches, nil); err == nil {
		return normalized[0], err
	}

	return nil, err
}

// GetTensorFromImageBuffer will take byte buffer and convert it into a tensor.
func GetTensorFromImageBuffer(imgBuffer bytes.Buffer, imgFormat string, numChan int64) (*tf.Tensor, error) {
	var err error
	var imgTensor *tf.Tensor

	imgTensor, err = tf.NewTensor(imgBuffer.String())
	if err != nil {
		return nil, err
	}

	var graph *tf.Graph
	var input, output tf.Output
	graph, input, output, err = createImageTransformGraph(imgFormat, numChan)
	if err != nil {
		return nil, err
	}

	var sess *tf.Session
	sess, err = tf.NewSession(graph, nil)
	if err != nil {
		return nil, err
	}

	defer sess.Close()

	feeds := map[tf.Output]*tf.Tensor{
		input: imgTensor,
	}

	fetches := []tf.Output{output}
	if normalized, err := sess.Run(feeds, fetches, nil); err == nil {
		return normalized[0], err
	}

	return nil, err
}

// createImageTransformGraph creates a TensorFlow graph that will perform normalization and image
// resizing on a given image.
func createImageTransformGraph(imgFormat string, numChan int64) (*tf.Graph, tf.Output, tf.Output, error) {
	s := op.NewScope()
	input := op.Placeholder(s, tf.String)

	// Decode either PNG or JPEG
	var decode tf.Output
	if imgFormat == "png" {
		decode = op.DecodePng(s, input, op.DecodePngChannels(numChan))
	} else {
		decode = op.DecodeJpeg(s, input, op.DecodeJpegChannels(numChan))
	}

	// Build the graph step by step
	expand := op.ExpandDims(s, op.Cast(s, decode, tf.Float), op.Const(s.SubScope("make_batch"), int32(0)))
	resize := op.ResizeBilinear(s, expand, op.Const(s.SubScope("size"), []int32{ImageHeight, ImageWidth}))
	subtractMean := op.Sub(s, resize, op.Const(s.SubScope("mean"), Mean))
	output := op.Div(s, subtractMean, op.Const(s.SubScope("scale"), Scale))

	// Graph is finalized!
	graph, err := s.Finalize()

	return graph, input, output, err
}
