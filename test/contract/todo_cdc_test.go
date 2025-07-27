package contract

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"todo-backend/internal/application/usecases"
	"todo-backend/internal/infrastructure/database"
	"todo-backend/internal/interfaces/handlers"
	"todo-backend/internal/interfaces/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TodoCDCProviderSuite implements Consumer Driven Contract provider tests
// These tests verify that our backend satisfies the exact contract expected by TodoFrontend
type TodoCDCProviderSuite struct {
	suite.Suite
	app *fiber.App
	db  *gorm.DB
}

func (suite *TodoCDCProviderSuite) SetupSuite() {
	// Setup in-memory database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)
	
	err = db.AutoMigrate(&database.SQLiteTodoModel{})
	suite.Require().NoError(err)
	
	suite.db = db
	
	// Setup application layers
	todoRepo := database.NewSQLiteTodoRepository(db)
	todoUseCase := usecases.NewTodoUseCase(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoUseCase)
	
	app := fiber.New()
	routes.SetupRoutes(app, todoHandler)
	suite.app = app
}

func (suite *TodoCDCProviderSuite) TearDownTest() {
	// Clean database between tests
	suite.db.Exec("DELETE FROM todos")
}

// CDC Test 1: "a request for all todos" with provider state "no todos exist"
// Expected: GET /api/todos → [] (empty array) with status 200
func (suite *TodoCDCProviderSuite) TestGetAllTodos_NoTodosExist() {
	// Provider State: no todos exist (database is clean)
	
	req := httptest.NewRequest("GET", "/api/todos", nil)
	resp, err := suite.app.Test(req)
	
	suite.NoError(err)
	
	// Contract assertion: Status 200
	suite.Equal(http.StatusOK, resp.StatusCode, "Contract violation: expected status 200")
	
	// Contract assertion: Content-Type header
	suite.Equal("application/json", resp.Header.Get("Content-Type"), 
		"Contract violation: expected Content-Type: application/json")
	
	// Contract assertion: Response body is empty array []
	var response interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	suite.NoError(err)
	
	// The response should be exactly [] (empty array)
	suite.Equal([]interface{}{}, response, 
		"Contract violation: expected empty array [] when no todos exist")
}

// CDC Test 2: "a request for all todos" with provider state "todos exist" 
// Expected: GET /api/todos → array with todo objects
func (suite *TodoCDCProviderSuite) TestGetAllTodos_TodosExist() {
	// Provider State: todos exist - seed test data exactly as in contract
	fixedTime := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	suite.db.Create(&database.SQLiteTodoModel{
		ID:        "uuid-123",
		Text:      "buy some milk",
		CreatedAt: fixedTime.Unix(),
		UpdatedAt: fixedTime.Unix(),
	})
	
	req := httptest.NewRequest("GET", "/api/todos", nil)
	resp, err := suite.app.Test(req)
	
	suite.NoError(err)
	
	// Contract assertion: Status 200
	suite.Equal(http.StatusOK, resp.StatusCode, "Contract violation: expected status 200")
	
	// Contract assertion: Content-Type header
	suite.Equal("application/json", resp.Header.Get("Content-Type"), 
		"Contract violation: expected Content-Type: application/json")
	
	// Contract assertion: Response body is array of todo objects
	var todos []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	suite.NoError(err)
	
	// Should have exactly one todo
	suite.Len(todos, 1, "Contract violation: expected exactly 1 todo")
	
	todo := todos[0]
	
	// Contract assertion: Todo object structure matches exactly
	suite.Equal("uuid-123", todo["id"], "Contract violation: id mismatch")
	suite.Equal("buy some milk", todo["text"], "Contract violation: text mismatch") 
	suite.Equal("2024-01-01T10:00:00.000Z", todo["createdAt"], 
		"Contract violation: createdAt format should be 2024-01-01T10:00:00.000Z")
	
	// Contract assertion: No extra fields (updatedAt should not be present)
	suite.Len(todo, 3, "Contract violation: todo should have exactly 3 fields (id, text, createdAt)")
	suite.Contains(todo, "id")
	suite.Contains(todo, "text") 
	suite.Contains(todo, "createdAt")
	suite.NotContains(todo, "updatedAt", "Contract violation: updatedAt should not be present")
}

// CDC Test 3: "a request to create a todo" with provider state "backend is ready to create todos"
// Expected: POST /api/todos with {"text": "buy some milk"} → created todo object with status 201
func (suite *TodoCDCProviderSuite) TestCreateTodo_BackendReady() {
	// Provider State: backend is ready to create todos (clean state)
	
	// Contract request body exactly as specified
	requestBody := map[string]string{
		"text": "buy some milk",
	}
	body, _ := json.Marshal(requestBody)
	
	req := httptest.NewRequest("POST", "/api/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := suite.app.Test(req)
	suite.NoError(err)
	
	// Contract assertion: Status 201 Created
	suite.Equal(http.StatusCreated, resp.StatusCode, 
		"Contract violation: expected status 201 Created")
	
	// Contract assertion: Content-Type header
	suite.Equal("application/json", resp.Header.Get("Content-Type"), 
		"Contract violation: expected Content-Type: application/json")
	
	// Contract assertion: Response body structure
	var todo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&todo)
	suite.NoError(err)
	
	// Contract assertions: Required fields present
	suite.Contains(todo, "id", "Contract violation: response missing 'id' field")
	suite.Contains(todo, "text", "Contract violation: response missing 'text' field")
	suite.Contains(todo, "createdAt", "Contract violation: response missing 'createdAt' field")
	
	// Contract assertions: Field values and formats
	suite.NotEmpty(todo["id"], "Contract violation: id should not be empty")
	suite.Equal("buy some milk", todo["text"], "Contract violation: text should match request")
	
	// Check createdAt format (should be ISO 8601 with .000Z)
	createdAt, ok := todo["createdAt"].(string)
	suite.True(ok, "Contract violation: createdAt should be string")
	suite.Regexp(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z$`, createdAt,
		"Contract violation: createdAt should match format YYYY-MM-DDTHH:mm:ss.sssZ")
	
	// Contract assertion: No extra fields (exactly 3 fields as per contract)
	suite.Len(todo, 3, "Contract violation: response should have exactly 3 fields (id, text, createdAt)")
	suite.NotContains(todo, "updatedAt", "Contract violation: updatedAt should not be present")
}

func TestTodoCDCProviderSuite(t *testing.T) {
	suite.Run(t, new(TodoCDCProviderSuite))
} 