package rest

import (
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

	handler.GET("/", r.getTodos)
	handler.PUT("/", r.createTodo)
	handler.PATCH("/:id/*action", r.updateTodo)
}

type getTodosParams struct {
	Offset int `form:"offset" json:"offset"`
	Size   int `form:"size" json:"size"`
}

func (r *todosRoutes) getTodos(ctx *gin.Context) {
	var params getTodosParams
	if err := ctx.BindQuery(&params); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	result, err := r.usecase.Get(ctx.Request.Context(), params.Offset, params.Size)
	if err != nil {
		errorResponse(ctx, http.StatusNotFound, "there is no todos")
		return
	}
	ctx.IndentedJSON(http.StatusOK, result)
}

func (r *todosRoutes) createTodo(ctx *gin.Context) {
	entity := &entity.ToDoItem{}
	if err := ctx.BindJSON(entity); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "wrong body data")
		return
	}
	id, _ := r.usecase.Create(ctx.Request.Context(), *entity)
	ctx.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

type updateAction uint

func (a updateAction) ToString() string {
	switch a {
	case done:
		return "/done"
	case delete:
		return "/delete"
	}
	return "unknown"
}

const (
	done updateAction = iota
	delete
)

func (r *todosRoutes) updateTodo(ctx *gin.Context) {
	action := ctx.Param("action")
	switch action {
	case done.ToString():
		r.markTodoDone(ctx)
	case delete.ToString():
		r.markTodoDeleted(ctx)
	default:
		errorResponse(ctx, http.StatusBadRequest, action+" action is not supported")
	}
}

func (r *todosRoutes) markTodoDeleted(ctx *gin.Context) {
	param, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	id := entity.ToDoItemId(param)
	err = r.usecase.MarkDeleted(ctx.Request.Context(), id)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "bad request")
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "deleted: " + strconv.FormatUint(param, 10),
	})
}

func (r *todosRoutes) markTodoDone(ctx *gin.Context) {
	param, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	id := entity.ToDoItemId(param)
	id, err = r.usecase.MarkDone(ctx.Request.Context(), id)
	if err != nil {
		errorResponse(ctx, http.StatusNotFound, "not found")
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "done: " + id.ToString(),
	})
}
