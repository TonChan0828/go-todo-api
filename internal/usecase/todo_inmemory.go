package usecase

import (
	"context"
	"sync"
	"time"

	"github.com/TonChan0828/go-todo-api/internal/domain"
	"github.com/google/uuid"
)

type InMemoryTodoUsecase struct {
	mu    sync.Mutex
	todos []*domain.Todo
}

func NewInMemoryTodoUsecase() *InMemoryTodoUsecase {
	return &InMemoryTodoUsecase{
		todos: make([]*domain.Todo, 0),
	}
}

func (u *InMemoryTodoUsecase) Create(ctx context.Context, title string) (*domain.Todo, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	now := time.Now()
	t := &domain.Todo{
		ID:        uuid.New(),
		Title:     title,
		CreatedAt: now,
		UpdatedAt: now,
	}

	u.todos = append(u.todos, t)

	return t, nil
}

func (u *InMemoryTodoUsecase) List(ctx context.Context) ([]*domain.Todo, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	// 呼び出し側が破壊的な変更をしても内部が壊れないようにコピー
	out := make([]*domain.Todo, len(u.todos))
	copy(out, u.todos)
	return out, nil
}
