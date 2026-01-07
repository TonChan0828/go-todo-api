package dto

type CreateTodoRequest struct {
	Title string `json:"title" example:"Buy milk"`
}

type UpdateTodoRequest struct {
	Completed bool `json:"completed" example:"true:"`
}
