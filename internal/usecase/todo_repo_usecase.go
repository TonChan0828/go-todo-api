package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/TonChan0828/go-todo-api/internal/domain"
	repository "github.com/TonChan0828/go-todo-api/internal/infrastructure"
	"github.com/google/uuid"
)

type RepoTodoUsecase struct {
	repo repository.TodoRepository
}

func NewRepoTodoUsecase(repo repository.TodoRepository) *RepoTodoUsecase {
	return &RepoTodoUsecase{repo: repo}
}

func (u *RepoTodoUsecase) Create(ctx context.Context, title string) (*domain.Todo, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	now := time.Now()
	t := &domain.Todo{ID: uuid.New(),
		Title:     title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := u.repo.Create(ctx, t); err != nil {
		return nil, err
	}
	return t, nil
}

func (u *RepoTodoUsecase) List(ctx context.Context) ([]*domain.Todo, error) {
	return u.repo.List(ctx)
}

func (u *RepoTodoUsecase) UpdateCompleted(ctx context.Context, id string, completed bool) (*domain.Todo, error) {
	if err := u.repo.UpdateCompleted(ctx, id, completed); err != nil {
		return nil, err
	}

	return u.repo.GetByID(ctx, id)
}

func (u *RepoTodoUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
