package main

import (
	"net/http"

	"github.com/TonChan0828/go-todo-api/internal/handler"
	"github.com/TonChan0828/go-todo-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	todoUC := usecase.NewInMemoryTodoUsecase()
	todoHandler := handler.NewTodoHandler(todoUC)
	handler.RegisterTodoRoutes(r, todoHandler)

	r.Run(":8080")
}
