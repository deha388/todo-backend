package routes

import (
	"todo-backend/internal/interfaces/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, todoHandler *handlers.TodoHandler) {
	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "Todo API is running",
		})
	})

	// API v1 routes
	api := app.Group("/api")
	
	// Todo routes - exactly as specified in requirements
	api.Get("/todos", todoHandler.GetTodos)        // GET /api/todos - List all todos
	api.Post("/todos", todoHandler.CreateTodo)     // POST /api/todos - Create new todo
} 