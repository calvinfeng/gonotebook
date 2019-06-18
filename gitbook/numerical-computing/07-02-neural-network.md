# Neural Networks

## Basics

What is a neural network? I believe you have seen a picture like this before, perhaps many times.

![multi-layer perceptons](../.gitbook/assets/neural_net.png)

Multi-layer perceptrons and neural networks are used in literature interchangeably. However, more concretely speaking, MLP is subset of neural networks. in general if you see a neural network that is feed-forward only, i.e. no cycles, then you are safe to call it a multi-layer perceptrons \(MLP\), otherwise they would be recurrent networks. For the purpose of introduction to neural networks, I will focus mainly on MLP. 

A neural network takes inputs and produces outputs. On the high level we can think of it as a function that takes inputs and computes outputs, using a set of hidden parameters. Let $$w$$ denote the hidden parameters, also known as weights. Let $$\vec{x}$$ denote our vector inputs and $$\vec{y}$$ denote our vector outputs. 

$$
f_w(\vec{x}) =\vec{y}
$$

In the simplest example, the neural network can be a linear function.

$$
f_w(\vec{x}) = W_{1}\vec{x} + W_{0}
$$

If input is a scalar, i.e. single dimensional vector, then we have the familiar equation of a line.

$$
f(x) = y = mx + b
$$

However, in practice, a neural network does not resemble a line because each layer has a non-linear activation function. Before we talk about non-linearity, let's focus a bit on the linear side first.

## Transformation

Recall that matrix multiplication represents a linear transformation of a vector. In case you need a little brush-up on linear algebra, you should take a look at [3Blue1Brown's lecture series](https://www.3blue1brown.com/essence-of-linear-algebra-page) on _Essence of Linear Algebra_. 

Suppose I want to apply a rotational transformation on a vector via my neural network, my $$W_1$$ would be a rotational matrix and $$W_0$$ is a zero matrix. Let me use the following notation to describe my transformation for better readability in my Python code.

$$
f(\vec{x}) = \vec{x}W_1 + W_0 \quad\text{where}\quad \vec{x} = \begin{vmatrix} 1 & 0 \end{vmatrix}
$$

Then

$$
W_1 = \begin{vmatrix} cos(\theta) & sin(\theta) \\ -sin(\theta) & cos(\theta)\end{vmatrix} \quad\text{and}\quad W_2 = \begin{vmatrix} 0 & 0 \\ 0 & 0 \end{vmatrix}
$$

I want to rotate my vector x by 90 degrees, which is $$\frac{\pi}{2}$$ in radians.

```python
import numpy as np

theta = np.pi / 2

W1 = np.array([
    [np.cos(theta), np.sin(theta)],
    [-np.sin(theta), np.cos(theta)],
])

W0 = np.array([
    [0.0, 0.0],
    [0.0, 0.0],
])

x = np.array([
    [1.0, 0.0],
])

y = np.dot(x, W1) + W0
```

Then the output will be

```text
[[0.0 1.0]]
```

I have rotated my horizontal vector 90 degrees and now it is a vertical vector on a Cartesian plane.

### Affine Transformation

The example above is commonly referred as affine transformation. The $$W_0$$ is known as offsets or **biases**. The generalized form would be describe by the following expression.

$$
\begin{vmatrix} x_0 & x_1 & ... & x_n \end{vmatrix} \begin{vmatrix}
W_{0,0} & W_{0,1} & ... & W_{0, m} \\
W_{1,0} & ... & ... & ... \\
... & ... & ... & ... \\
W_{n, 0} & ... & ... & W_{n, m}
 \end{vmatrix}
 + \begin{vmatrix}
b_{0} & b_{1} & ... & b_{m} 
\end{vmatrix}
$$

### Nonlinear Transformation

The power of neural network lies in the fact that it can model any function, linear or nonlinear. So far we have only applied affine transformation, we are missing an ingredient to allow our network to model nonlinear functions. The key ingredient we need here is a nonlinear transformation, also commonly known as nonlinear activation.



