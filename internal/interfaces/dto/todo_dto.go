package dto

import (
	"todo-backend/internal/domain/entities"
)

type CreateTodoRequest struct {
	Text string `json:"text" validate:"required,min=1,max=500"`
}

// Removed: TodoResponse and TodoListResponse structs
// These are replaced by ContractTodoResponse in contract_dto.go for better contract compliance

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

// Removed: ToTodoResponse and ToTodoListResponse functions
// These are replaced by ToContractTodoResponse and ToContractTodoList in contract_dto.go

func (req *CreateTodoRequest) ToEntity() *entities.Todo {
	return entities.NewTodo(req.Text)
}

func SuccessResponse(data interface{}, message string) APIResponse {
	return APIResponse{
		Success: true,
		Data:    data,
		Message: message,
	}
}

func ErrorResponse(err string) APIResponse {
	return APIResponse{
		Success: false,
		Error:   err,
	}
} 