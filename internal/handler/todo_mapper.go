package handler

import (
	"github.com/TonChan0828/go-todo-api/internal/domain"
	"github.com/TonChan0828/go-todo-api/internal/handler/dto"
)

func toTOdoResponse(t *domain.Todo) dto.TodoResponse {
	return dto.TodoResponse{
		ID:        t.ID.String(),
		Title:     t.Title,
		Completed: t.Completed,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
