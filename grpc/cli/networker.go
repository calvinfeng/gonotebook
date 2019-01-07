package cli

import (
	"bytes"
	"context"
	"fmt"
	"go-academy/grpc/pb/todo"
	"net/http"
	"time"
)

type Networker interface {
	Get(id int) ([]byte, error)
	Set(id int, data []byte) error
}

func NewHTTPNetworker(endpoint string) Networker {
	return &HTTPNetworker{
		client:   http.DefaultClient,
		endpoint: endpoint,
	}
}

type HTTPNetworker struct {
	client   *http.Client
	endpoint string
}

func (net *HTTPNetworker) Get(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/%d/", net.endpoint, id)

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

func (net *HTTPNetworker) Set(id int, data []byte) error {
	url := fmt.Sprintf("%s/%d/", net.endpoint, id)

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

func NewGRPCNetworker() Networker {
	return &GRPCNetworker{}
}

type GRPCNetworker struct {
	client todo.TodoClient
}

func (net *GRPCNetworker) Get(id int) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := net.client.Get(ctx, &todo.TodoRequest{Id: 1})
	if err != nil {
		return nil, err
	}

	fmt.Println(res)

	return nil, nil
}

func (net *GRPCNetworker) Set(id int, data []byte) error {
	return nil
}
