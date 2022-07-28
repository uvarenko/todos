package postgre

import (
	"fmt"
	"os"

	"github.com/uvarenko/todo-list/internal/entity"
	"github.com/uvarenko/todo-list/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(l *logger.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_DB_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatal("can't connect to db", err)
	}
	if err = db.AutoMigrate(&entity.ToDoItem{}); err != nil {
		l.Fatal("can't migrate db", err)
	}

	return db
}
