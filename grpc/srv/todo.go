package srv

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-academy/grpc/pb/planner"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRPCTodoServer returns a RPC server.
func NewRPCTodoServer() *RPCTodoServer {
	return &RPCTodoServer{}
}

// RPCTodoServer implements TodoServer interface.
type RPCTodoServer struct{}

// Get handles RPC request for getting a Todo resource.
func (srv *RPCTodoServer) Get(ctx context.Context, req *planner.TodoRequest) (*planner.TodoResponse, error) {
	model, err := store.Get(req.Id)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	return &planner.TodoResponse{Data: data}, nil
}

// Set handles RPC request for setting a Todo resource.
func (srv *RPCTodoServer) Set(ctx context.Context, req *planner.TodoRequest) (*planner.TodoResponse, error) {
	if req.Data == nil {
		return nil, errors.New("please provide data")
	}

	model := &TodoModel{}
	err := json.Unmarshal(req.Data.GetValue(), model)
	if err != nil {
		return nil, err
	}

	model.ID = req.Id

	err = store.Set(model)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	return &planner.TodoResponse{Data: data}, nil
}

// NewHTTPTodoServer returns a HTTP server.
func NewHTTPTodoServer() *HTTPTodoServer {
	return &HTTPTodoServer{
		&http.Server{},
	}
}

// HTTPTodoServer is a wrapper around standard http server.
type HTTPTodoServer struct {
	*http.Server
}

// Register configures routes for HTTP server.
func (s *HTTPTodoServer) Register() {
	mux := mux.NewRouter().StrictSlash(true)
	mux.Handle("/todos/{id}/", http.HandlerFunc(s.handleGet)).Methods("GET")
	mux.Handle("/todos/{id}/", http.HandlerFunc(s.handlePost)).Methods("POST")
	s.Handler = mux
}

func (s *HTTPTodoServer) handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting something")
}

func (s *HTTPTodoServer) handlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create something")
}
