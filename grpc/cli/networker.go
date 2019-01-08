package cli

import (
	"bytes"
	"context"
	"fmt"
	"go-academy/grpc/pb/planner"
	"net/http"
	"time"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"

	"google.golang.org/grpc"
)

type Networker interface {
	Get(id int64) ([]byte, error)
	Set(id int64, data []byte) ([]byte, error)
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

func (net *HTTPNetworker) Get(id int64) ([]byte, error) {
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

func (net *HTTPNetworker) Set(id int64, data []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%d/", net.endpoint, id)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
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

func NewGRPCNetworker(conn *grpc.ClientConn) Networker {
	return &GRPCNetworker{planner.NewTodoClient(conn)}
}

type GRPCNetworker struct {
	client planner.TodoClient
}

func (net *GRPCNetworker) Get(id int64) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &planner.TodoRequest{
		Id: &wrappers.Int64Value{Value: id},
	}

	res, err := net.client.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (net *GRPCNetworker) Set(id int64, data []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &planner.TodoRequest{
		Id:   &wrappers.Int64Value{Value: id},
		Data: &wrappers.BytesValue{Value: data},
	}

	res, err := net.client.Set(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
