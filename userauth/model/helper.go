package model

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	byteArray := make([]byte, n)
	_, err := rand.Read(byteArray)
	if err != nil {
		return nil, err
	}

	return byteArray, nil
}

func GenerateRandomString(length int) (string, error) {
	bytes, err := GenerateRandomBytes(length)
	return base64.URLEncoding.EncodeToString(bytes), err
}
