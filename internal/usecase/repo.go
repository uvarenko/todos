package usecase

import (
	"context"

	"github.com/uvarenko/todo-list/internal/entity"
)

type ToDoRepo interface {
	Add(context.Context, entity.ToDoItem) (uint, error)
	MarkDone(context.Context, entity.ToDoItemId) (entity.ToDoItemId, error)
	MarkDeleted(context.Context, entity.ToDoItemId) error
	Get(context.Context, int, int) ([]entity.ToDoItem, error)
}
