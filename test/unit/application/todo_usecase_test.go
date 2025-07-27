package application

import (
	"context"
	"errors"
	"testing"
	"todo-backend/internal/application/usecases"
	"todo-backend/internal/domain/entities"
	"todo-backend/internal/domain/repositories"
	"todo-backend/internal/interfaces/dto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTodoRepository for application layer testing
type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) Create(ctx context.Context, todo *entities.Todo) (*entities.Todo, error) {
	args := m.Called(ctx, todo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Todo), args.Error(1)
}

func (m *MockTodoRepository) GetAll(ctx context.Context) ([]*entities.Todo, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Todo), args.Error(1)
}

func (m *MockTodoRepository) GetByID(ctx context.Context, id string) (*entities.Todo, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Todo), args.Error(1)
}

// Application Layer Use Case Tests
// These test BUSINESS LOGIC ORCHESTRATION only

func TestTodoUseCase_CreateTodo_Success(t *testing.T) {
	// Given
	mockRepo := &MockTodoRepository{}
	useCase := usecases.NewTodoUseCase(mockRepo)
	ctx := context.Background()
	
	req := dto.CreateTodoRequest{Text: "Test Todo"}
	expectedTodo := entities.NewTodo("Test Todo")
	
	mockRepo.On("Create", ctx, mock.AnythingOfType("*entities.Todo")).Return(expectedTodo, nil)
	
	// When
	result, err := useCase.CreateTodo(ctx, req)
	
	// Then
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedTodo.ID, result.ID)
	assert.Equal(t, "Test Todo", result.Text)
	
	mockRepo.AssertExpectations(t)
}

func TestTodoUseCase_CreateTodo_EmptyText_ShouldFail(t *testing.T) {
	// Given
	mockRepo := &MockTodoRepository{}
	useCase := usecases.NewTodoUseCase(mockRepo)
	ctx := context.Background()
	
	req := dto.CreateTodoRequest{Text: ""}
	
	// When
	result, err := useCase.CreateTodo(ctx, req)
	
	// Then
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "cannot be empty")
	
	// Repository should not be called
	mockRepo.AssertNotCalled(t, "Create")
}

func TestTodoUseCase_CreateTodo_RepositoryError(t *testing.T) {
	// Given
	mockRepo := &MockTodoRepository{}
	useCase := usecases.NewTodoUseCase(mockRepo)
	ctx := context.Background()
	
	req := dto.CreateTodoRequest{Text: "Test Todo"}
	repoError := errors.New("database connection failed")
	
	mockRepo.On("Create", ctx, mock.AnythingOfType("*entities.Todo")).Return(nil, repoError)
	
	// When
	result, err := useCase.CreateTodo(ctx, req)
	
	// Then
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to create todo")
	assert.Contains(t, err.Error(), "database connection failed")
	
	mockRepo.AssertExpectations(t)
}

func TestTodoUseCase_GetAllTodos_Success(t *testing.T) {
	// Given
	mockRepo := &MockTodoRepository{}
	useCase := usecases.NewTodoUseCase(mockRepo)
	ctx := context.Background()
	
	expectedTodos := []*entities.Todo{
		entities.NewTodo("Todo 1"),
		entities.NewTodo("Todo 2"),
	}
	
	mockRepo.On("GetAll", ctx).Return(expectedTodos, nil)
	
	// When
	result, err := useCase.GetAllTodos(ctx)
	
	// Then
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	
	assert.Equal(t, expectedTodos[0].ID, result[0].ID)
	assert.Equal(t, "Todo 1", result[0].Text)
	assert.Equal(t, expectedTodos[1].ID, result[1].ID)
	assert.Equal(t, "Todo 2", result[1].Text)
	
	mockRepo.AssertExpectations(t)
}

func TestTodoUseCase_GetAllTodos_RepositoryError(t *testing.T) {
	// Given
	mockRepo := &MockTodoRepository{}
	useCase := usecases.NewTodoUseCase(mockRepo)
	ctx := context.Background()
	
	repoError := errors.New("connection timeout")
	mockRepo.On("GetAll", ctx).Return(nil, repoError)
	
	// When
	result, err := useCase.GetAllTodos(ctx)
	
	// Then
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to get todos")
	assert.Contains(t, err.Error(), "connection timeout")
	
	mockRepo.AssertExpectations(t)
}

func TestTodoUseCase_GetTodoByID_Success(t *testing.T) {
	// Given
	mockRepo := &MockTodoRepository{}
	useCase := usecases.NewTodoUseCase(mockRepo)
	ctx := context.Background()
	
	todoID := "test-id-123"
	expectedTodo := entities.NewTodo("Test Todo")
	expectedTodo.ID = todoID
	
	mockRepo.On("GetByID", ctx, todoID).Return(expectedTodo, nil)
	
	// When
	result, err := useCase.GetTodoByID(ctx, todoID)
	
	// Then
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, todoID, result.ID)
	assert.Equal(t, "Test Todo", result.Text)
	
	mockRepo.AssertExpectations(t)
}

func TestTodoUseCase_GetTodoByID_NotFound(t *testing.T) {
	// Given
	mockRepo := &MockTodoRepository{}
	useCase := usecases.NewTodoUseCase(mockRepo)
	ctx := context.Background()
	
	todoID := "non-existent-id"
	mockRepo.On("GetByID", ctx, todoID).Return(nil, repositories.ErrTodoNotFound)
	
	// When
	result, err := useCase.GetTodoByID(ctx, todoID)
	
	// Then
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
	assert.Contains(t, err.Error(), todoID)
	
	mockRepo.AssertExpectations(t)
} 