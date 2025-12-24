package handler

import (
	"github.com/TonChan0828/go-todo-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	uc usecase.TodoUsecase
}

func NewTodoHandler(uc usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{uc: uc}
}

func (h *TodoHandler) Create(c *gin.Context) {
	panic("not implemented")
}

func (h *TodoHandler) List(c *gin.Context) {
	panic("not implemented")
}
