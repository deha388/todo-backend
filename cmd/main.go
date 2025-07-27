package main

import (
	"errors"
	"log"
	"todo-backend/internal/application/usecases"
	"todo-backend/internal/infrastructure/config"
	"todo-backend/internal/infrastructure/database"
	"todo-backend/internal/interfaces/handlers"
	"todo-backend/internal/interfaces/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Printf("⚠️  Failed to load configuration: %v", err)
		log.Println("📝 Using default configuration...")
		cfg = &config.Config{
			Server: config.ServerConfig{
				Host: "0.0.0.0",
				Port: 8083,
			},
			Database: config.DatabaseConfig{
				Type: "sqlite",
				File: "todo.db",
			},
		}
	} else {
		log.Printf("Configuration loaded from configs/config.yaml")
	}

	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to SQLite database: %v", err)
	}

	todoRepo := database.NewSQLiteTodoRepository(db)
	todoUseCase := usecases.NewTodoUseCase(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoUseCase)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		},
	})

	routes.SetupRoutes(app, todoHandler)
	log.Println("✅ Routes configured")

	log.Println("\n📋 Available Endpoints:")
	log.Println("  GET    /health           - Health check")
	log.Println("  GET    /api/todos        - List all todos")
	log.Println("  POST   /api/todos        - Create new todo")

	serverAddr := cfg.GetServerAddress()
	log.Printf("\n🌐 Server starting on %s", serverAddr)
	log.Printf("🎯 API Base URL: http://%s/api", serverAddr)

	if err := app.Listen(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
