package grpc

import (
	"context"

	"github.com/fuku01/test-v2-api/pkg/grpc/pb"
	"github.com/fuku01/test-v2-api/pkg/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoHandler interface {
	GetTodoList(context.Context, *emptypb.Empty) (*pb.TodoList, error)
}

type todoHandler struct {
	tu usecase.TodoUsecase
}

func NewTodoHandler(tu usecase.TodoUsecase) TodoHandler {
	return &todoHandler{tu: tu}
}

func (th *todoHandler) GetTodoList(ctx context.Context, empty *emptypb.Empty) (*pb.TodoList, error) {
	_, err := th.tu.GetTodoList()
	if err != nil {
		return nil, err
	}

	return &pb.TodoList{}, nil
}
