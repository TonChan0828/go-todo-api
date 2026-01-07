package dto

import "time"

type TodoResponse struct {
	ID        string    `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Title     string    `json:"title" example:"Buy milk"`
	Completed bool      `json:"completed" example:"false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TodoListResponse struct {
	Items []TodoResponse `json:"items"`
}
