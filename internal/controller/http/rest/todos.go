package rest

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uvarenko/todo-list/internal/entity"
	"github.com/uvarenko/todo-list/internal/usecase"
)

type todosRoutes struct {
	usecase *usecase.ToDoUseCase
}

func newTodosRoutes(handler *gin.RouterGroup, usecase *usecase.ToDoUseCase) {
	r := &todosRoutes{usecase: usecase}

	handler.GET("/", validateGetTodos(), r.getTodos)
	handler.PUT("/", validateCreateTodo(), r.createTodo)
	handler.PATCH("/:id/*action", validateUpdateTodo(), r.updateTodo)
}

func (r *todosRoutes) getTodos(ctx *gin.Context) {
	var params getTodosParams
	ctx.BindQuery(&params)
	result, err := r.usecase.Get(ctx.Request.Context(), params.Offset, params.Size)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, errors.New("there is no todos"))
		return
	}
	ctx.IndentedJSON(http.StatusOK, result)
}

func (r *todosRoutes) createTodo(ctx *gin.Context) {
	var entity entity.ToDoItem
	id, _ := r.usecase.Create(ctx.Request.Context(), entity)
	ctx.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

func (r *todosRoutes) updateTodo(ctx *gin.Context) {
	action := updateAction(ctx.Param("action"))
	switch action {
	case done:
		r.markTodoDone(ctx)
	case delete:
		r.markTodoDeleted(ctx)
	}
}

func (r *todosRoutes) markTodoDeleted(ctx *gin.Context) {
	param, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	id := entity.ToDoItemId(param)
	err := r.usecase.MarkDeleted(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "deleted: " + strconv.FormatUint(param, 10),
	})
}

func (r *todosRoutes) markTodoDone(ctx *gin.Context) {
	param, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	id := entity.ToDoItemId(param)
	id, err := r.usecase.MarkDone(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "done: " + id.ToString(),
	})
}
