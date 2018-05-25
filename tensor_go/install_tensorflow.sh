#!/usr/bin/env bash
TF_TYPE="cpu"
TARGET_DIRECTORY='/usr/local'
curl -L \
    "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-${TF_TYPE}-$(go env GOOS)-x86_64-1.6.0.tar.gz" |
sudo tar -C $TARGET_DIRECTORY -xz