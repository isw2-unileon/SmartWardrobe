package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"backend/internal/dto"
	"backend/internal/handlers"
)

type MockUserService struct {
	MockToken string
	MockErr   error
}

func (m *MockUserService) Login(userName string, password string) (string, error) {
	return m.MockToken, m.MockErr
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func TestUserHandler_Login_Success(t *testing.T) {

	mockService := &MockUserService{
		MockToken: "token-jwt-falso",
		MockErr:   nil,
	}
	handler := handlers.NewUserHandler(mockService)

	router := setupRouter()
	router.POST("/api/login", handler.Login)

	loginData := dto.LoginDto{
		Email:    "testuser@test.com",
		Password: "123456",
	}
	jsonValue, err := json.Marshal(loginData)
	if err != nil {
		t.Fatalf("Error al serializar el JSON: %v", err)
	}

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("We were expecting status 200 OK, but we got: %d. Body: %s", w.Code, w.Body.String())
	}

	if !bytes.Contains(w.Body.Bytes(), []byte("token-jwt-falso")) {
		t.Errorf("The response does not contain the expected token. Actual response: %s", w.Body.String())
	}
}

func TestUserHandler_Login_InvalidJSON(t *testing.T) {
	mockService := &MockUserService{}
	handler := handlers.NewUserHandler(mockService)

	router := setupRouter()
	router.POST("/api/login", handler.Login)

	invalidJson := []byte(`{"userName": "testuser", password: "123"}`)

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(invalidJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("We were expecting status 400 BAD REQUEST, but we got: %d", w.Code)
	}
}

func TestUserHandler_Login_Unauthorized(t *testing.T) {

	mockService := &MockUserService{
		MockToken: "",
		MockErr:   errors.New("credenciales inválidas"),
	}
	handler := handlers.NewUserHandler(mockService)

	router := setupRouter()
	router.POST("/api/login", handler.Login)

	loginData := dto.LoginDto{
		Email:    "hacker",
		Password: "wrongpassword",
	}
	jsonValue, err := json.Marshal(loginData)
	if err != nil {
		t.Fatalf("Error al serializar el JSON: %v", err)
	}

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("We were expecting status 401 UNAUTHORIZED, but we got: %d", w.Code)
	}
}
