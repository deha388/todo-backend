package domain

import (
	"testing"
	"todo-backend/internal/domain/entities"

	"github.com/stretchr/testify/assert"
)

// Domain Unit Tests
// These test ONLY business logic without external dependencies

func TestTodo_Creation(t *testing.T) {
	t.Run("should create todo with valid text", func(t *testing.T) {
		// Given
		text := "Learn Clean Architecture with TDD"

		// When  
		todo := entities.NewTodo(text)

		// Then
		assert.NotEmpty(t, todo.ID, "Todo ID should not be empty")
		assert.Equal(t, text, todo.Text, "Todo text should match input")
		assert.NotZero(t, todo.CreatedAt, "CreatedAt should be set")
		assert.NotZero(t, todo.UpdatedAt, "UpdatedAt should be set")
	})
} 