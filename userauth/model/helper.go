package model

import (
	"crypto/rand"
	"encoding/base64"
)

func generateRandomBytes(n int) ([]byte, error) {
	byteArray := make([]byte, n)
	_, err := rand.Read(byteArray)
	if err != nil {
		return nil, err
	}

	return byteArray, nil
}

func generateRandomString(length int) (string, error) {
	bytes, err := generateRandomBytes(length)
	return base64.URLEncoding.EncodeToString(bytes), err
}
