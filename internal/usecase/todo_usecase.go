package usecase

import (
	"context"

	"github.com/TonChan0828/go-todo-api/internal/domain"
)

type TodoUsecase interface {
	Create(ctx context.Context, title string) (*domain.Todo, error)
	List(ctx context.Context) ([]*domain.Todo, error)
	UpdateCompleted(ctx context.Context, id string, completed bool) (*domain.Todo, error)
	Delete(ctx context.Context, id string) error
}
