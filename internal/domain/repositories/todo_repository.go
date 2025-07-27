package repositories

import (
	"context"
	"errors"
	"todo-backend/internal/domain/entities"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
	ErrTodoExists   = errors.New("todo already exists")
)

type TodoRepository interface {
	Create(ctx context.Context, todo *entities.Todo) (*entities.Todo, error)
	
	GetAll(ctx context.Context) ([]*entities.Todo, error)
	
	GetByID(ctx context.Context, id string) (*entities.Todo, error)
} 