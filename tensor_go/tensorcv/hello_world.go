package tensorcv

import (
	"fmt"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

// HelloWorldFromTF will initialize a session and run it. It will print Hello to the screen.
func HelloWorldFromTF() {
	// Construct a graph with an operation that produces a string constant
	s := op.NewScope()
	c := op.Const(s, "Hello from TensorFlow version "+tf.Version())

	graph, err := s.Finalize()
	if err != nil {
		panic(err)
	}

	// Execute the graph in a session
	sess, err := tf.NewSession(graph, nil)
	if err != nil {
		panic(err)
	}

	defer sess.Close()

	output, err := sess.Run(nil, []tf.Output{c}, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(output[0].Value())
}
