package handler

import (
	"net/http"

	"github.com/TonChan0828/go-todo-api/internal/handler/dto"
	"github.com/TonChan0828/go-todo-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	uc usecase.TodoUsecase
}

func NewTodoHandler(uc usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{uc: uc}
}

// CreateTodo godoc
// @Summary      Create a todo
// @Description Create a new todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo  body  dto.CreateTodoRequest  true  "Todo request"
// @Success      201   {object}  dto.TodoResponse
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /todos [post]
func (h *TodoHandler) Create(c *gin.Context) {
	var req dto.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.uc.Create(c.Request.Context(), req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.TodoResponse{
		ID:        todo.ID.String(),
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	c.JSON(http.StatusCreated, res)
}

// ListTodos godoc
// @Summary      List todos
// @Description Get all todos
// @Tags         todos
// @Produce      json
// @Success      200  {object}  dto.TodoListResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /todos [get]
func (h *TodoHandler) List(c *gin.Context) {
	todos, err := h.uc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// UpdateTodo godoc
// @Summary      Update todo completion
// @Description Update completed status of a todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id    path   string                 true  "Todo ID"
// @Param        todo  body   dto.UpdateTodoRequest  true  "Update todo"
// @Success      200   {object}  dto.TodoResponse
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      404   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /todos/{id} [put]
func (h *TodoHandler) UpdateCompleted(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		Completed bool `json:"completed"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.uc.UpdateCompleted(c.Request.Context(), id, req.Completed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo godoc
// @Summary      Delete a todo
// @Description Delete a todo by ID
// @Tags         todos
// @Param        id  path  string  true  "Todo ID"
// @Success      204
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /todos/{id} [delete]
func (h *TodoHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.uc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
