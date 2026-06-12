package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockUserService struct {
	createUserFn  func(user *User) error
	getUsersFn    func() ([]User, error)
	getUserByIDFn func(userID uint) (*User, error)
	updateUserFn  func(userID uint, updatedData *User) (*User, error)
	deleteUserFn  func(userID uint) error
}

func (m *mockUserService) GetUsers() ([]User, error) {
	return m.getUsersFn()
}

func (m *mockUserService) GetUserByID(userID uint) (*User, error) {
	return m.getUserByIDFn(userID)
}

func (m *mockUserService) CreateUser(user *User) error {
	return m.createUserFn(user)
}

func (m *mockUserService) UpdateUser(userID uint, updatedData *User) (*User, error) {
	return m.updateUserFn(userID, updatedData)
}

func (m *mockUserService) DeleteUser(userID uint) error {
	return m.deleteUserFn(userID)
}

func TestCreateUser_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`{"username": 123}`))
	c.Request.Header.Set("Content-Type", "application/json")

	h := &handlerUsers{service: &mockUserService{}}
	h.CreateUser(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d but got %d", http.StatusBadRequest, w.Code)
	}
}

func TestCreateUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var createdUser *User
	mockService := &mockUserService{
		createUserFn: func(user *User) error {
			createdUser = user
			return nil
		},
	}

	requestBody := map[string]string{
		"username": "testuser",
		"email":    "test@example.com",
		"password": "secret123",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("error generating request body: %v", err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	h := &handlerUsers{service: mockService}
	h.CreateUser(c)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status %d but got %d", http.StatusCreated, w.Code)
	}

	var responseBody map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if responseBody["username"] != "testuser" {
		t.Fatalf("expected username %q but got %q", "testuser", responseBody["username"])
	}

	if responseBody["email"] != "test@example.com" {
		t.Fatalf("expected email %q but got %q", "test@example.com", responseBody["email"])
	}

	if _, exists := responseBody["password"]; exists {
		t.Fatal("password field should not be included in the response")
	}

	if createdUser == nil || createdUser.Username != "testuser" {
		t.Fatal("expected service CreateUser to receive the parsed user data")
	}
}
