package handlers

import (
	"todo-backend/internal/application/usecases"
	"todo-backend/internal/interfaces/dto"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	todoUseCase *usecases.TodoUseCase
}

// NewTodoHandler creates a new TodoHandler
func NewTodoHandler(todoUseCase *usecases.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		todoUseCase: todoUseCase,
	}
}

// GetTodos handles GET /api/todos
func (h *TodoHandler) GetTodos(c *fiber.Ctx) error {
	ctx := c.Context()

	// Get all todos
	todos, err := h.todoUseCase.GetAllTodos(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.ErrorResponse(err.Error()),
		)
	}

	// Return contract-compliant response (plain array)
	contractResponse := dto.ToContractTodoList(todos)
	return c.Status(fiber.StatusOK).JSON(contractResponse)
}

// CreateTodo handles POST /api/todos
func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	ctx := c.Context()

	// Parse request body
	var req dto.CreateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.ErrorResponse("Invalid request body"),
		)
	}

	//validate
	if req.Text == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.ErrorResponse("Text field is required"),
		)
	}

	//create
	todo, err := h.todoUseCase.CreateTodo(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.ErrorResponse(err.Error()),
		)
	}

	// Return created todo in contract format
	contractResponse := dto.ToContractTodoResponse(todo)
	return c.Status(fiber.StatusCreated).JSON(contractResponse)
}
