# TensorFlow in Go
## Install TensorFlow Go-binding
Enter the following command to your terminal
```
TF_TYPE="cpu" # Change to "gpu" for GPU support
TARGET_DIRECTORY='/usr/local'
curl -L \
   "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-${TF_TYPE}-$(go env GOOS)-x86_64-1.1.0.tar.gz" |

sudo tar -C $TARGET_DIRECTORY -xz
sudo ldconfig
```

Or run the shell script
```
./install_tensorflow.sh
```

And run
```
sudo ldconfig
```

And then you are ready to run `dep ensure`!