package main

// @title Go Todo API
// @version 1.0
// @description This is a simple todo API server built with Go and Gin.
// @host localhost:8080
// @BasePath /v1

import (
	"log"
	"net/http"

	v1handler "github.com/TonChan0828/go-todo-api/internal/handler/v1"
	"github.com/TonChan0828/go-todo-api/internal/infrastructure/db"
	"github.com/TonChan0828/go-todo-api/internal/repository"
	"github.com/TonChan0828/go-todo-api/internal/usecase"
	"github.com/gin-gonic/gin"

	_ "github.com/TonChan0828/go-todo-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// todoUC := usecase.NewInMemoryTodoUsecase()

	sqlDB, err := db.OpenPostgres()
	if err != nil {
		log.Fatal(err)
	}
	q := db.New(sqlDB)
	todoRepo := repository.NewPostgresTodoRepository(q)
	todoUC := usecase.NewRepoTodoUsecase(todoRepo)

	// ======v1 Group ======
	v1 := r.Group("/v1")
	todoHandler := v1handler.NewTodoHandler(todoUC)
	v1handler.RegisterTodoRoutes(v1, todoHandler)

	r.Run(":8080")
}
