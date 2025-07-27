package integration

import (
	"context"
	"testing"
	"todo-backend/internal/domain/entities"
	"todo-backend/internal/domain/repositories"
	"todo-backend/internal/infrastructure/database"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Repository Integration Tests
// These test the actual SQLite repository implementation with real database

func setupRepositoryTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	
	err = db.AutoMigrate(&database.SQLiteTodoModel{})
	require.NoError(t, err)
	
	return db
}

func TestSQLiteTodoRepository_Create_Integration(t *testing.T) {
	t.Run("should create todo successfully", func(t *testing.T) {
		// Given
		db := setupRepositoryTestDB(t)
		repo := database.NewSQLiteTodoRepository(db)
		ctx := context.Background()
		
		todo := entities.NewTodo("Test todo for repository")
		
		// When
		result, err := repo.Create(ctx, todo)
		
		// Then
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, todo.ID, result.ID)
		assert.Equal(t, "Test todo for repository", result.Text)
		
		// Verify persisted in database
		var count int64
		db.Model(&database.SQLiteTodoModel{}).Count(&count)
		assert.Equal(t, int64(1), count)
		
		// Verify database content
		var model database.SQLiteTodoModel
		db.First(&model)
		assert.Equal(t, todo.ID, model.ID)
		assert.Equal(t, "Test todo for repository", model.Text)
	})
}

func TestSQLiteTodoRepository_GetAll_Integration(t *testing.T) {
	t.Run("should return all todos ordered by created_at DESC", func(t *testing.T) {
		// Given
		db := setupRepositoryTestDB(t)
		repo := database.NewSQLiteTodoRepository(db)
		ctx := context.Background()
		
		// Create test data directly in database to control timestamps
		db.Create(&database.SQLiteTodoModel{
			ID:        "todo-1",
			Text:      "First todo",
			CreatedAt: 1000,
		})
		db.Create(&database.SQLiteTodoModel{
			ID:        "todo-2",
			Text:      "Second todo",
			CreatedAt: 2000,
		})
		
		// When
		todos, err := repo.GetAll(ctx)
		
		// Then
		assert.NoError(t, err)
		assert.Len(t, todos, 2)
		
		// Should be ordered by created_at DESC (newest first)
		assert.Equal(t, "todo-2", todos[0].ID)
		assert.Equal(t, "Second todo", todos[0].Text)
		assert.Equal(t, "todo-1", todos[1].ID)
		assert.Equal(t, "First todo", todos[1].Text)
	})
	
	t.Run("should return empty list when no todos exist", func(t *testing.T) {
		// Given
		db := setupRepositoryTestDB(t)
		repo := database.NewSQLiteTodoRepository(db)
		ctx := context.Background()
		
		// When
		todos, err := repo.GetAll(ctx)
		
		// Then
		assert.NoError(t, err)
		assert.Len(t, todos, 0)
	})
}

func TestSQLiteTodoRepository_GetByID_Integration(t *testing.T) {
	t.Run("should return todo by ID", func(t *testing.T) {
		// Given
		db := setupRepositoryTestDB(t)
		repo := database.NewSQLiteTodoRepository(db)
		ctx := context.Background()
		
		// Seed test data
		db.Create(&database.SQLiteTodoModel{
			ID:        "test-id-123",
			Text:      "Find me",
			CreatedAt: 1000,
		})
		
		// When
		todo, err := repo.GetByID(ctx, "test-id-123")
		
		// Then
		assert.NoError(t, err)
		assert.NotNil(t, todo)
		assert.Equal(t, "test-id-123", todo.ID)
		assert.Equal(t, "Find me", todo.Text)
	})
	
	t.Run("should return ErrTodoNotFound when todo does not exist", func(t *testing.T) {
		// Given
		db := setupRepositoryTestDB(t)
		repo := database.NewSQLiteTodoRepository(db)
		ctx := context.Background()
		
		// When
		todo, err := repo.GetByID(ctx, "non-existent-id")
		
		// Then
		assert.Error(t, err)
		assert.Nil(t, todo)
		assert.Equal(t, repositories.ErrTodoNotFound, err)
	})
}

// Entity-Model Conversion Integration Tests
func TestSQLiteTodoModel_ToEntity_Integration(t *testing.T) {
	t.Run("should convert model to entity correctly", func(t *testing.T) {
		// Given
		model := database.SQLiteTodoModel{
			ID:        "test-id",
			Text:      "Test todo",
			CreatedAt: 1609459200, // 2021-01-01 00:00:00 UTC
			UpdatedAt: 1609459260, // 2021-01-01 00:01:00 UTC
		}
		
		// When
		entity, err := model.ToEntity()
		
		// Then
		assert.NoError(t, err)
		assert.Equal(t, "test-id", entity.ID)
		assert.Equal(t, "Test todo", entity.Text)
		assert.Equal(t, int64(1609459200), entity.CreatedAt.Unix())
		assert.Equal(t, int64(1609459260), entity.UpdatedAt.Unix())
	})
}

func TestSQLiteTodoModel_FromEntity_Integration(t *testing.T) {
	t.Run("should convert entity to model correctly", func(t *testing.T) {
		// Given
		entity := entities.NewTodo("Test todo")
		
		// When
		var model database.SQLiteTodoModel
		model.FromEntity(entity)
		
		// Then
		assert.Equal(t, entity.ID, model.ID)
		assert.Equal(t, "Test todo", model.Text)
		assert.Equal(t, entity.CreatedAt.Unix(), model.CreatedAt)
		assert.Equal(t, entity.UpdatedAt.Unix(), model.UpdatedAt)
	})
} 