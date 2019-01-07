package main

import (
	"bytes"
	"fmt"
	"net/http"
)

type Networker interface {
	Get(url string) ([]byte, error)
	Set(url string, data []byte) error
}

func NewHTTPNetworker() Networker {
	return &HTTPNetworker{
		client: http.DefaultClient,
	}
}

type HTTPNetworker struct {
	client *http.Client
}

func (net *HTTPNetworker) Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token1234567890")

	res, err := net.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	buf := bytes.NewBuffer([]byte{})
	buf.ReadFrom(res.Body)

	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("bad status %d - %s", res.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}

func (net *HTTPNetworker) Set(url string, data []byte) error {
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token1234567890")

	res, err := net.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	buf := bytes.NewBuffer([]byte{})
	buf.ReadFrom(res.Body)

	if res.StatusCode >= 300 {
		return fmt.Errorf("bad status %d - %s", res.StatusCode, buf.String())
	}

	return nil
}
