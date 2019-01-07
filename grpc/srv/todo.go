package srv

import (
	"context"
	"go-academy/grpc/pb/todo"
)

func NewRPCTodoServer() *RPCTodoServer {
	return &RPCTodoServer{}
}

type RPCTodoServer struct {
}

func (srv *RPCTodoServer) Get(ctx context.Context, req *todo.TodoRequest) (*todo.TodoResponse, error) {
	return nil, nil
}
