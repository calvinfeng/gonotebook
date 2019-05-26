package model

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// ValidationError is an error that captures invalid data.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (err *ValidationError) Error() string {
	return fmt.Sprintf("ValidationError %s: %s", err.Field, err.Message)
}

func generateRandomBytes(n int) ([]byte, error) {
	byteArray := make([]byte, n)
	_, err := rand.Read(byteArray)
	if err != nil {
		return nil, err
	}

	return byteArray, nil
}

func fakeJWTToken(length int) (string, error) {
	bytes, err := generateRandomBytes(length)
	return base64.URLEncoding.EncodeToString(bytes), err
}
