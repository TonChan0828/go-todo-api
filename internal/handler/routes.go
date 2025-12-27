package handler

import "github.com/gin-gonic/gin"

func RegisterTodoRoutes(r *gin.Engine, h *TodoHandler) {
	todos := r.Group("/todos")
	{
		todos.POST("", h.Create)
		todos.GET("", h.List)
		todos.PUT("/:id", h.UpdateCompleted)
		todos.DELETE("/:id", h.Delete)
	}
}
