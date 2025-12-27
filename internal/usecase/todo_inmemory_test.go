package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestInMemoryTodoUsecase_CreateAndList(t *testing.T) {
	uc := NewInMemoryTodoUsecase()
	ctx := context.Background()

	// Create
	todo, err := uc.Create(ctx, "test task")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if todo.Title != "test task" {
		t.Errorf("expected title 'test task', got '%s'", todo.Title)
	}
	if todo.ID == uuid.Nil {
		t.Errorf("expected non-nil UUID")
	}

	// List
	todos, err := uc.List(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(todos))
	}
	if todos[0].Title != "test task" {
		t.Errorf("unexpected todo title")
	}
}

func TestInMemoryTodoUsecase_UpdateAndDelete(t *testing.T) {
	uc := NewInMemoryTodoUsecase()
	ctx := context.Background()

	todo, _ := uc.Create(ctx, "task")
	id := todo.ID.String()

	// Update
	updated, err := uc.UpdateCompleted(ctx, id, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !updated.Completed {
		t.Errorf("expected completed = true")
	}

	// Delete
	if err := uc.Delete(ctx, id); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	todos, _ := uc.List(ctx)
	if len(todos) != 0 {
		t.Errorf("expected 0 todos after delete")
	}
}
