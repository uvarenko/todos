package repo

import (
	"context"

	"github.com/uvarenko/todo-list/internal/entity"
	"github.com/uvarenko/todo-list/internal/usecase"
	"gorm.io/gorm"
)

type toDoRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) usecase.ToDoRepo {
	return &toDoRepo{
		db: db,
	}
}

func (r *toDoRepo) Add(c context.Context, item entity.ToDoItem) (uint, error) {
	result := r.db.Create(&item)
	inserted := result.Statement.Model.(*entity.ToDoItem)
	return uint(inserted.Id), result.Error
}

func (r *toDoRepo) MarkDone(c context.Context, id entity.ToDoItemId) (entity.ToDoItemId, error) {
	result := r.db.Model(&entity.ToDoItem{}).Where("id=?", id).Update("is_done", true)
	return id, result.Error
}

func (r *toDoRepo) MarkDeleted(c context.Context, id entity.ToDoItemId) error {
	result := r.db.Model(&entity.ToDoItem{}).Where("id=?", id).Update("is_deleted", true)
	return result.Error
}

func (r *toDoRepo) Get(c context.Context, offset int, size int) ([]entity.ToDoItem, error) {
	items := make([]entity.ToDoItem, size)
	result := r.db.Limit(size).Offset(offset).Find(&items)
	return items, result.Error
}
