package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler holds the route handlers for the Todo API.
type Handler struct {
	store *Store
}

// NewHandler creates a Handler backed by the provided store.
func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes attaches the Todo CRUD routes to the given router group.
func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/todos", h.Create)
	rg.GET("/todos", h.List)
	rg.PUT("/todos/:id", h.Update)
	rg.DELETE("/todos/:id", h.Delete)
}

type createRequest struct {
	Title string `json:"title" binding:"required"`
}

// Create handles POST /todos.
func (h *Handler) Create(c *gin.Context) {
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := h.store.Add(req.Title)
	c.JSON(http.StatusCreated, t)
}

// List handles GET /todos.
func (h *Handler) List(c *gin.Context) {
	todos := h.store.List()
	if todos == nil {
		todos = []Todo{}
	}
	c.JSON(http.StatusOK, todos)
}

type updateRequest struct {
	Title     string `json:"title"     binding:"required"`
	Completed bool   `json:"completed"`
}

// Update handles PUT /todos/:id.
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req updateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t, ok := h.store.Update(id, req.Title, req.Completed)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	c.JSON(http.StatusOK, t)
}

// Delete handles DELETE /todos/:id.
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if !h.store.Delete(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
