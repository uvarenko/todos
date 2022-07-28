package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/uvarenko/todo-list/internal/usecase"
)

func NewRouter(handler *gin.Engine, usecase *usecase.ToDoUseCase) {
	// add middlewares
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// setup routes
	h := handler.Group("/todo")
	{
		newTodosRoutes(h, usecase)
	}
}
