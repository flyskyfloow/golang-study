package todo_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flyskyfloow/golang-study/todo"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	store := todo.NewStore()
	h := todo.NewHandler(store)
	r := gin.New()
	v1 := r.Group("/api/v1")
	h.RegisterRoutes(v1)
	return r
}

func TestCreateTodo(t *testing.T) {
	r := setupRouter()

	body := `{"title":"Buy milk"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/todos", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	var got todo.Todo
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if got.ID == 0 {
		t.Error("expected non-zero ID")
	}
	if got.Title != "Buy milk" {
		t.Errorf("expected title 'Buy milk', got %q", got.Title)
	}
	if got.Completed {
		t.Error("expected completed=false on creation")
	}
}

func TestCreateTodo_MissingTitle(t *testing.T) {
	r := setupRouter()

	body := `{}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/todos", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestListTodos_Empty(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/todos", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var got []todo.Todo
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected empty list, got %d items", len(got))
	}
}

func TestListTodos(t *testing.T) {
	r := setupRouter()

	// Create two todos first
	for _, title := range []string{"Task A", "Task B"} {
		body := `{"title":"` + title + `"}`
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/todos", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(httptest.NewRecorder(), req)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/todos", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var got []todo.Todo
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("expected 2 todos, got %d", len(got))
	}
}

func TestUpdateTodo(t *testing.T) {
	r := setupRouter()

	// Create a todo
	createBody := `{"title":"Original"}`
	createReq, _ := http.NewRequest(http.MethodPost, "/api/v1/todos", bytes.NewBufferString(createBody))
	createReq.Header.Set("Content-Type", "application/json")
	createW := httptest.NewRecorder()
	r.ServeHTTP(createW, createReq)

	var created todo.Todo
	json.NewDecoder(createW.Body).Decode(&created)

	// Update it
	updateBody := `{"title":"Updated","completed":true}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/todos/1", bytes.NewBufferString(updateBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var got todo.Todo
	json.NewDecoder(w.Body).Decode(&got)
	if got.Title != "Updated" {
		t.Errorf("expected title 'Updated', got %q", got.Title)
	}
	if !got.Completed {
		t.Error("expected completed=true")
	}
}

func TestUpdateTodo_NotFound(t *testing.T) {
	r := setupRouter()

	body := `{"title":"Ghost","completed":false}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/todos/999", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestDeleteTodo(t *testing.T) {
	r := setupRouter()

	// Create a todo
	body := `{"title":"To delete"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/todos", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(httptest.NewRecorder(), req)

	// Delete it
	w := httptest.NewRecorder()
	delReq, _ := http.NewRequest(http.MethodDelete, "/api/v1/todos/1", nil)
	r.ServeHTTP(w, delReq)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", w.Code)
	}

	// Verify it's gone
	listW := httptest.NewRecorder()
	listReq, _ := http.NewRequest(http.MethodGet, "/api/v1/todos", nil)
	r.ServeHTTP(listW, listReq)

	var todos []todo.Todo
	json.NewDecoder(listW.Body).Decode(&todos)
	if len(todos) != 0 {
		t.Errorf("expected empty list after delete, got %d items", len(todos))
	}
}

func TestDeleteTodo_NotFound(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/todos/999", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestUpdateTodo_InvalidID(t *testing.T) {
	r := setupRouter()

	body := `{"title":"x","completed":false}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/todos/abc", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestDeleteTodo_InvalidID(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/todos/abc", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
