package entity

import "strconv"

type ToDoItemId uint

type ToDoItem struct {
	Id        ToDoItemId `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	Title     string     `gorm:"size:255;not null;" json:"title" binding:"required"`
	Content   string     `gorm:"size:255;not null;" json:"content" binding:"required"`
	Created   string     `json:"created_time"`
	Notify    string     `json:"notify_time"`
	IsDone    bool       `json:"is_done"`
	IsDeleted bool       `json:"is_deleted"`
}

func (id ToDoItemId) ToString() string {
	return strconv.FormatUint(uint64(id), 10)
}
