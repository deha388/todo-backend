package database

import (
	"context"
	"fmt"
	"todo-backend/internal/domain/entities"
	"todo-backend/internal/domain/repositories"

	"gorm.io/gorm"
)

// SQLiteTodoRepository implements TodoRepository using SQLite
type SQLiteTodoRepository struct {
	db *gorm.DB
}

// NewSQLiteTodoRepository creates a new SQLite todo repository
func NewSQLiteTodoRepository(db *gorm.DB) repositories.TodoRepository {
	return &SQLiteTodoRepository{
		db: db,
	}
}

// SQLiteTodoModel represents the database model for SQLite todos
type SQLiteTodoModel struct {
	ID        string `gorm:"primaryKey;type:text"`
	Text      string `gorm:"not null;type:text"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
}

// TableName returns the table name for SQLiteTodoModel
func (SQLiteTodoModel) TableName() string {
	return "todos"
}

// ToEntity converts SQLiteTodoModel to domain entity
func (tm *SQLiteTodoModel) ToEntity() (*entities.Todo, error) {
	todo := &entities.Todo{
		ID:        tm.ID,
		Text:      tm.Text,
		CreatedAt: timeFromUnix(tm.CreatedAt),
		UpdatedAt: timeFromUnix(tm.UpdatedAt),
	}

	return todo, nil
}

// FromEntity converts domain entity to SQLiteTodoModel
func (tm *SQLiteTodoModel) FromEntity(todo *entities.Todo) {
	tm.ID = todo.ID
	tm.Text = todo.Text
	tm.CreatedAt = todo.CreatedAt.Unix()
	tm.UpdatedAt = todo.UpdatedAt.Unix()
}

// Create creates a new todo
func (r *SQLiteTodoRepository) Create(ctx context.Context, todo *entities.Todo) (*entities.Todo, error) {
	model := &SQLiteTodoModel{}
	model.FromEntity(todo)

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	return model.ToEntity()
}

// GetAll retrieves all todos
func (r *SQLiteTodoRepository) GetAll(ctx context.Context) ([]*entities.Todo, error) {
	var models []SQLiteTodoModel
	if err := r.db.WithContext(ctx).Order("created_at DESC").Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get todos: %w", err)
	}

	todos := make([]*entities.Todo, len(models))
	for i, model := range models {
		todo, err := model.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("failed to convert todo model: %w", err)
		}
		todos[i] = todo
	}

	return todos, nil
}

// GetByID retrieves a todo by its ID
func (r *SQLiteTodoRepository) GetByID(ctx context.Context, id string) (*entities.Todo, error) {
	var model SQLiteTodoModel
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, repositories.ErrTodoNotFound
		}
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	return model.ToEntity()
}
