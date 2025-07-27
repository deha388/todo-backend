package usecases

import (
	"context"
	"errors"
	"fmt"
	"todo-backend/internal/domain/entities"
	"todo-backend/internal/domain/repositories"
	"todo-backend/internal/interfaces/dto"
)

// test-driven - no code

type TodoUseCase struct {
	todoRepo repositories.TodoRepository
}

func NewTodoUseCase(todoRepo repositories.TodoRepository) *TodoUseCase {
	return &TodoUseCase{
		todoRepo: todoRepo,
	}
}

func (uc *TodoUseCase) CreateTodo(ctx context.Context, req dto.CreateTodoRequest) (*entities.Todo, error) {

	if req.Text == "" {
		return nil, fmt.Errorf("todo text cannot be empty")
	}

	todo := req.ToEntity()
	created, err := uc.todoRepo.Create(ctx, todo)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	return created, nil
}

func (uc *TodoUseCase) GetAllTodos(ctx context.Context) ([]*entities.Todo, error) {

	todos, err := uc.todoRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get todos: %w", err)
	}

	return todos, nil
}

func (uc *TodoUseCase) GetTodoByID(ctx context.Context, id string) (*entities.Todo, error) {

	if id == "" {
		return nil, fmt.Errorf("todo ID cannot be empty")
	}

	todo, err := uc.todoRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repositories.ErrTodoNotFound) {
			return nil, fmt.Errorf("todo with ID %s not found", id)
		}
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	return todo, nil
}
