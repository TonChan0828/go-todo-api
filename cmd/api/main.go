package main

import (
	"net/http"

	"github.com/TonChan0828/go-todo-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	todoHandler := handler.NewTodoHandler(nil)
	handler.RegisterTodoRoutes(r, todoHandler)

	r.Run(":8080")
}
