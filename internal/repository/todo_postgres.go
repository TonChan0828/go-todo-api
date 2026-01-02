package repository

import (
	"context"

	"github.com/TonChan0828/go-todo-api/internal/domain"
	"github.com/TonChan0828/go-todo-api/internal/infrastructure/db"
	"github.com/google/uuid"
)

type PostgresTodoRepository struct {
	q *db.Queries
}

func NewPostgresTodoRepository(q *db.Queries) *PostgresTodoRepository {
	return &PostgresTodoRepository{q: q}
}

func (r *PostgresTodoRepository) Create(ctx context.Context, todo *domain.Todo) error {
	return r.q.CreateTodo(ctx, db.CreateTodoParams{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	})
}

func (r *PostgresTodoRepository) List(ctx context.Context) ([]*domain.Todo, error) {
	rows, err := r.q.ListTodos((ctx))
	if err != nil {
		return nil, err
	}

	todos := make([]*domain.Todo, 0, len(rows))
	for _, r := range rows {
		todos = append(todos, &domain.Todo{
			ID:        r.ID,
			Title:     r.Title,
			Completed: r.Completed,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}
	return todos, nil
}

func (r *PostgresTodoRepository) UpdateCompleted(ctx context.Context, id string, completed bool) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.q.UpdateTodoCompleted(ctx, db.UpdateTodoCompletedParams{
		ID:        uid,
		Completed: completed,
	})
}

func (r *PostgresTodoRepository) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.q.DeleteTodo(ctx, uid)
}

func (r *PostgresTodoRepository) GetByID(ctx context.Context, id string) (*domain.Todo, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	row, err := r.q.GetTodoByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return &domain.Todo{
		ID:        row.ID,
		Title:     row.Title,
		Completed: row.Completed,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}, nil
}
