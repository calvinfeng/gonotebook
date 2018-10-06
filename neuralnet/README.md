# Neural Network
## Project Requirements
If you are familiar with vectorized implementation of neural network, feel free to jump ahead
and start watching videos on Golang part, otherwise we need to go over some the basics and 
mathematics of neural nets.

## Jupyter Notebook
I will use Python to teach math because I can write LaTex in Jupyter notebooks. Also, `numpy` is 
very convenient for matrix operations. We will see that we have a `numpy` equivalent in Golang 
called `gonum` (I wonder why not `numgo`?)

So let's get started by installing `pip`. I think `easy_install` is provided by Mac OS X, so we 
don't need to use Homebrew.
```
sudo easy_install pip
```

Once you have `pip`, now use it to install `virtualenv` (Python virtual environment)
```
pip install virtualenv
```

If permission denied, use `sudo`. This is equivalent to install `npm` globally. Remember that `pip`
is a package manager for Python, just like npm for Node.
```
sudo pip install virtualenv
```

Now go to `neural_net` directory and create a virtual environment
```
cd $GOPATH/src/go-academy/neural_net/
virtualenv environment
```

Activate your environment
```
source environment/bin/activate
```

Install all the required dependencies
```
pip install numpy
pip install matplotlib
pip install jupyter
```

Now you are good to go, let's run Jupyter!!! Make a directory called `notebooks` in `neural_net`
```
mkdir notebooks
cd notebooks
jupyter notebook
```

## Videos
Need more time to work on the videos for this project...