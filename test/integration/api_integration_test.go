package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-backend/internal/application/usecases"
	"todo-backend/internal/infrastructure/database"
	"todo-backend/internal/interfaces/handlers"
	"todo-backend/internal/interfaces/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// APIIntegrationTestSuite tests the integration between HTTP layer and application layer
// This follows the TDD workflow: integration test → routing code → business code
type APIIntegrationTestSuite struct {
	suite.Suite
	app *fiber.App
	db  *gorm.DB
}

func (suite *APIIntegrationTestSuite) SetupSuite() {
	// Setup test database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)
	
	err = db.AutoMigrate(&database.SQLiteTodoModel{})
	suite.Require().NoError(err)
	
	suite.db = db
	
	// Setup application (integration level)
	todoRepo := database.NewSQLiteTodoRepository(db)
	todoUseCase := usecases.NewTodoUseCase(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoUseCase)
	
	app := fiber.New()
	routes.SetupRoutes(app, todoHandler)
	suite.app = app
}

func (suite *APIIntegrationTestSuite) TearDownTest() {
	suite.db.Exec("DELETE FROM todos")
}

// Integration Test: HTTP → Handler → UseCase → Repository → Database
func (suite *APIIntegrationTestSuite) TestCreateTodoAPI_Integration() {
	// Test the complete request flow
	reqBody := map[string]string{"text": "Integration test todo"}
	body, _ := json.Marshal(reqBody)
	
	req := httptest.NewRequest("POST", "/api/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := suite.app.Test(req)
	suite.NoError(err)
	suite.Equal(http.StatusCreated, resp.StatusCode)
	
	// Verify response structure (contract-compliant format)
	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)
	
	suite.NotEmpty(response["id"])
	suite.Equal("Integration test todo", response["text"])
	suite.NotEmpty(response["createdAt"])
	// Note: updatedAt is no longer returned per contract requirements
	suite.NotContains(response, "updatedAt")
	
	// Verify data persisted in database
	var count int64
	suite.db.Model(&database.SQLiteTodoModel{}).Count(&count)
	suite.Equal(int64(1), count)
}

func (suite *APIIntegrationTestSuite) TestGetTodosAPI_Integration() {
	// Seed test data
	suite.db.Create(&database.SQLiteTodoModel{
		ID:   "test-id-1",
		Text: "Test Todo 1",
	})
	suite.db.Create(&database.SQLiteTodoModel{
		ID:   "test-id-2", 
		Text: "Test Todo 2",
	})
	
	req := httptest.NewRequest("GET", "/api/todos", nil)
	resp, err := suite.app.Test(req)
	
	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
	
	var todos []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&todos)
	
	suite.Len(todos, 2)
	
	// Verify response format (contract-compliant format)
	for _, todo := range todos {
		suite.NotEmpty(todo["id"])
		suite.NotEmpty(todo["text"])
		suite.NotEmpty(todo["createdAt"])
		// Note: updatedAt is no longer returned per contract requirements
		suite.NotContains(todo, "updatedAt")
	}
}

// Error handling integration tests
func (suite *APIIntegrationTestSuite) TestErrorHandling_BadJSON() {
	req := httptest.NewRequest("POST", "/api/todos", bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := suite.app.Test(req)
	suite.NoError(err)
	suite.Equal(http.StatusBadRequest, resp.StatusCode)
}

func TestAPIIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(APIIntegrationTestSuite))
} 