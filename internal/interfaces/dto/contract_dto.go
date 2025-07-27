package dto

import (
	"time"
	"todo-backend/internal/domain/entities"
)

// ContractTodoResponse represents the exact response format expected by the frontend contract
// Based on the Pact contract: only id, text, and createdAt fields in UTC format
type ContractTodoResponse struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"` // ISO 8601 format with .000Z
}

// ToContractTodoResponse converts entity to contract-compliant response
func ToContractTodoResponse(todo *entities.Todo) ContractTodoResponse {
	return ContractTodoResponse{
		ID:        todo.ID,
		Text:      todo.Text,
		CreatedAt: formatTimeForContract(todo.CreatedAt),
	}
}

// ToContractTodoList converts entity slice to contract-compliant response array
func ToContractTodoList(todos []*entities.Todo) []ContractTodoResponse {
	if len(todos) == 0 {
		return []ContractTodoResponse{} // Return empty array, not nil
	}
	
	responses := make([]ContractTodoResponse, len(todos))
	for i, todo := range todos {
		responses[i] = ToContractTodoResponse(todo)
	}
	return responses
}

// formatTimeForContract converts time.Time to the exact format expected by contract
// Format: "2024-01-01T10:00:00.000Z" (ISO 8601 with milliseconds and Z suffix)
func formatTimeForContract(t time.Time) string {
	// Ensure time is in UTC and format with .000Z suffix
	utc := t.UTC()
	return utc.Format("2006-01-02T15:04:05.000Z")
} 