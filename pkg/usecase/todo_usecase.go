package usecase

import (
	"github.com/fuku01/test-v2-grpc/pkg/domain/entity"
	"github.com/fuku01/test-v2-grpc/pkg/domain/repository"
)

type TodoUsecase interface {
	GetTodoList() ([]entity.Todo, error)
}

type todoUsecase struct {
	tr repository.TodoRepository
}

func NewTodoUsecase(tr repository.TodoRepository) TodoUsecase {
	return &todoUsecase{tr: tr}
}

func (tu *todoUsecase) GetTodoList() ([]entity.Todo, error) {
	return tu.tr.GetTodoList()
}
