from tensorflow.python.tools import inspect_checkpoint as ckpt
from keras.applications.resnet50 import ResNet50
from keras import backend as K

import tensorflow as tf
import time


def export_from_keras():
    """
    Loads ResNet50 model from Keras and then uses Keras backend session to find the graph and save 
    it to a protobuf file. However, in order to access variables and operations of the graph, one 
    must know the correct naming scheme for every operations. In this case, ResNet50 has a variable 
    named "input_1" for image tensor input and a "fc1000/softmax" as the final layer of the network.
    """
    model = ResNet50(weights='imagenet')
    print 'Model is loaded!'
    sess = K.get_session()

    operation_names = []
    for op in sess.graph.get_operations():
        operation_names.append(str(op.name))

    print 'List of first few operations', operation_names[0:10]

    builder = tf.saved_model.builder.SavedModelBuilder('resnet')
    builder.add_meta_graph_and_variables(K.get_session(), ['serve'])
    builder.save()

    print 'Model is saved!'


def export_from_checkpoint():
    dir = './resnet50/resnet_v2_50.ckpt'
    ckpt.print_tensors_in_checkpoint_file(dir, tensor_name='', 
                                               all_tensors=False, 
                                               all_tensor_names=False)
    tf.reset_default_graph()
    saver = tf.train.Saver()
    with tf.Session() as sess:
        saver.restore(sess, './resnet50/resnet_v2_50.ckpt')
        print 'Model is restored!'

        builder = tf.saved_model.builder.SavedModelBuilder('resnet_50_model')
        builder.add_meta_graph_and_variables(sess, ['serve'])
        builder.save()


def main():
    export_from_keras()


if __name__ == '__main__':
    main()