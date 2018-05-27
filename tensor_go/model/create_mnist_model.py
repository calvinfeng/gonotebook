from tensorflow.examples.tutorials.mnist import input_data
from matplotlib import pyplot as plt

import tensorflow as tf


DATA_DIR = 'datasets/'
def main():
    """
    Visit https://www.tensorflow.org/versions/r1.2/get_started/mnist/beginners and download the
    MNIST data
    """
    mnist = input_data.read_data_sets(DATA_DIR, one_hot=True)

    # Assuming each image is 28x28 with 1 channel
    x = tf.placeholder(tf.float32, [None, 28, 28, 1], name='image_input')
    W = tf.Variable(tf.zeros([784, 10]))
    b = tf.Variable(tf.zeros([10]))
    y = tf.add(tf.matmul(tf.reshape(x, [-1, 784]), W), b)
    labels = tf.placeholder(tf.float32, [None, 10])
    cross_entropy_loss = tf.reduce_mean(tf.nn.softmax_cross_entropy_with_logits(labels=labels, 
                                                                                logits=y))
    train_step = tf.train.GradientDescentOptimizer(0.5).minimize(cross_entropy_loss)

    with tf.Session() as sess:
        with tf.device("/cpu:0"):
            sess.run(tf.global_variables_initializer())
            for i in range(1000):
                # The data come in as (N, 784), thus we must reshape it.
                batch_x, batch_label = mnist.train.next_batch(100)
                batch_x = batch_x.reshape((-1, 28, 28, 1))

                loss, _ = sess.run([cross_entropy_loss, train_step], 
                                   feed_dict={x: batch_x, labels: batch_label})
                print 'Iteration %d: %f' % (i + 1, loss)

            infer = tf.argmax(y, axis=1, name='infer')
            truth = tf.argmax(labels, axis=1)
            correct_prediction = tf.equal(infer, truth)
            accuracy = tf.reduce_mean(tf.cast(correct_prediction, tf.float32))

            test_x = mnist.test.images.reshape((-1, 28, 28, 1))
            test_y = mnist.test.labels
            print 'Accuracy on test data:', sess.run(accuracy, 
                                                     feed_dict={x: test_x, labels: test_y})

            print 'Time to save the graph!'
            builder = tf.saved_model.builder.SavedModelBuilder('model')
            builder.add_meta_graph_and_variables(sess, ['serve'])
            builder.save()


if __name__ == '__main__':
    main()