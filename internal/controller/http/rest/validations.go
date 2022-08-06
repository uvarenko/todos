package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uvarenko/todo-list/internal/entity"
)

func validateGetTodos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var query getTodosParams
		err := ctx.BindQuery(&query)
		if query.Offset < 0 {
			err = fmt.Errorf("offset should be >= 0, %w", err)
		}
		if query.Size <= 0 {
			err = fmt.Errorf("size should be > 0, %w", err)
		}
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Next()
	}
}

func validateUpdateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var responseErr error
		action := updateAction(ctx.Param("action"))
		if action != done && action != delete {
			responseErr = fmt.Errorf("action should be in range [done, delete], %w", responseErr)
		}
		_, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
		if err != nil {
			responseErr = fmt.Errorf("id param should be >= 0, %w", responseErr)
		}
		if responseErr != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": responseErr.Error(),
			})
			return
		}

		ctx.Next()
	}
}

func validateCreateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var entity entity.ToDoItem
		if err := ctx.BindJSON(&entity); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Next()
	}
}
