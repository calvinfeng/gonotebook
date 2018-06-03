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
	ImageHeight = 224
	ImageWidth  = 224
	Mean        = float32(0)
	Scale       = float32(1)
)

// GetTensorFromImagePath creates a tensor by taking an image path.
func GetTensorFromImagePath(imgPath, imgFormat string, numChan int64) (*tf.Tensor, error) {
	var err error

	imgFile, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}

	defer imgFile.Close()

	var imgBuffer bytes.Buffer
	io.Copy(&imgBuffer, imgFile)

	return GetTensorFromImageBuffer(imgBuffer, imgFormat, numChan)
}

// GetTensorFromImageBuffer creates a tensor by taking a byte buffer.
func GetTensorFromImageBuffer(imgBuffer bytes.Buffer, imgFormat string,
	numChan int64) (*tf.Tensor, error) {

	imgTensor, err := tf.NewTensor(imgBuffer.String())
	if err != nil {
		return nil, err
	}

	graph, input, output, err := createImageTransformGraph(imgFormat, numChan)
	if err != nil {
		return nil, err
	}

	sess, err := tf.NewSession(graph, nil)
	if err != nil {
		return nil, err
	}

	defer sess.Close()

	feeds := map[tf.Output]*tf.Tensor{
		input: imgTensor,
	}

	fetches := []tf.Output{output}
	normalized, err := sess.Run(feeds, fetches, nil)
	if err == nil {
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

// Class refers to the classification of an image
type Class struct {
	Prob  float32
	Index int
}

// Sort will perform quick sort in place on a list of Class elements.
func Sort(list []Class, low, high int) {
	if low < high {
		partIdx := Partition(list, low, high)
		Sort(list, low, partIdx-1)
		Sort(list, partIdx+1, high)
	}
}

// Partition will perform quick sort partitioning in place.
func Partition(list []Class, low, high int) int {
	pivot := list[high]

	i := low - 1 // Index of the smaller element
	for j := low; j < high; j++ {
		if list[j].Prob <= pivot.Prob {
			i++
			temp := list[i]
			list[i] = list[j]
			list[j] = temp
		}
	}

	temp := list[i+1]
	list[i+1] = list[high]
	list[high] = temp

	return i + 1
}
