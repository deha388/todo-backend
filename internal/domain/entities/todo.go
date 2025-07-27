package entities

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTodo(text string) *Todo {
	now := time.Now()
	return &Todo{
		ID:        uuid.New().String(),
		Text:      text,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
