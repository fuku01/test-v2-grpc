package repository

import (
	"github.com/fuku01/test-v2-api/pkg/domain/entity"
	"github.com/fuku01/test-v2-api/pkg/domain/repository"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &todoRepository{db: db}
}

func (tr *todoRepository) GetTodoList() ([]entity.Todo, error) {
	return nil, nil
}
