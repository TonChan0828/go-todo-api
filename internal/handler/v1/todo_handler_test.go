package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TonChan0828/go-todo-api/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// テスト用スタブ
type stubTodoUsecase struct{}

func (s *stubTodoUsecase) Create(ctx context.Context, title string) (*domain.Todo, error) {
	return &domain.Todo{
		ID:        uuid.New(),
		Title:     title,
		Completed: false,
	}, nil
}

func (s *stubTodoUsecase) List(ctx context.Context) ([]*domain.Todo, error) {
	return []*domain.Todo{}, nil
}

func (s *stubTodoUsecase) UpdateCompleted(ctx context.Context, id string, conpleted bool) (*domain.Todo, error) {
	return nil, nil
}

func (s *stubTodoUsecase) Delete(ctx context.Context, id string) error {
	return nil
}

// テスト
func TestTodoHandler_Create(t *testing.T) {
	gin.SetMode((gin.TestMode))
	uc := &stubTodoUsecase{}
	h := NewTodoHandler(uc)

	r := gin.New()
	r.POST("/todos", h.Create)

	body := map[string]string{"title": "test todo"}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}
}
