package usecase

import (
	"context"

	"github.com/uvarenko/todo-list/internal/entity"
)

type ToDoUseCase struct {
	repo ToDoRepo
}

func NewToDoUseCase(repo ToDoRepo) *ToDoUseCase {
	return &ToDoUseCase{
		repo: repo,
	}
}

func (u *ToDoUseCase) Get(c context.Context, offset int, size int) ([]entity.ToDoItem, error) {
	return u.repo.Get(c, offset, size)
}

func (u *ToDoUseCase) Create(c context.Context, data entity.ToDoItem) (uint, error) {
	return u.repo.Add(c, data)
}

func (u *ToDoUseCase) MarkDeleted(c context.Context, id entity.ToDoItemId) error {
	return u.repo.MarkDeleted(c, id)
}

func (u *ToDoUseCase) MarkDone(c context.Context, id entity.ToDoItemId) (entity.ToDoItemId, error) {
	return u.repo.MarkDone(c, id)
}
