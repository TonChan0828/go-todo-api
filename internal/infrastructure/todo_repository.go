package repository

import (
	"context"

	"github.com/TonChan0828/go-todo-api/internal/domain"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *domain.Todo) error
	List(ctx context.Context) ([]*domain.Todo, error)
	GetByID(ctx context.Context, id string) (*domain.Todo, error)
	UpdateCompleted(ctx context.Context, id string, completed bool) error
	Delete(ctx context.Context, id string) error
}
