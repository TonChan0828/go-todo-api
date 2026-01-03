package repository

import (
	"context"
	"testing"
	"time"

	"github.com/TonChan0828/go-todo-api/internal/domain"
	"github.com/TonChan0828/go-todo-api/internal/infrastructure/db"
	"github.com/google/uuid"
)

func TestPostgresTodoRepository_CreateAndGetByID(t *testing.T) {
	tx, cleanup := setupTestTx(t)
	defer cleanup()

	q := db.New(tx)
	repo := NewPostgresTodoRepository(q)

	now := time.Now()
	todo := &domain.Todo{
		ID:        uuid.New(),
		Title:     "db test",
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Create
	if err := repo.Create(context.Background(), todo); err != nil {
		t.Fatalf("create failed: %v", err)
	}

	// GetByID
	got, err := repo.GetByID(context.Background(), todo.ID.String())
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}

	if got.Title != todo.Title {
		t.Fatalf("expected title %q, got %q", todo.Title, got.Title)
	}
	if got.Completed != todo.Completed {
		t.Fatalf("completed mismatch")
	}
}
