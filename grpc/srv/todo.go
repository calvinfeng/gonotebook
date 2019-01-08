package srv

import (
	"context"
	"encoding/json"
	"errors"
	"go-academy/grpc/pb/planner"
)

func NewRPCTodoServer() planner.TodoServer {
	return &RPCTodoServer{}
}

type RPCTodoServer struct{}

func (srv *RPCTodoServer) Get(ctx context.Context, req *planner.TodoRequest) (*planner.TodoResponse, error) {
	if req.Id == nil {
		return nil, errors.New("please provide ID")
	}

	model, err := store.Get(req.Id.GetValue())
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	return &planner.TodoResponse{Data: data}, nil
}

func (srv *RPCTodoServer) Set(ctx context.Context, req *planner.TodoRequest) (*planner.TodoResponse, error) {
	if req.Data == nil {
		return nil, errors.New("please provide data")
	}

	model := &TodoModel{}
	err := json.Unmarshal(req.Data.GetValue(), model)
	if err != nil {
		return nil, err
	}

	if req.Id != nil {
		model.ID = req.Id.GetValue()
	}

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
