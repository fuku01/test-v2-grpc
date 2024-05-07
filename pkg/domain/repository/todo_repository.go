package repository

import "github.com/fuku01/test-v2-grpc/pkg/domain/entity"

type TodoRepository interface {
	GetTodoList() ([]entity.Todo, error)
}
