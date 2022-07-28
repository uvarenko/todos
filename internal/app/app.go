package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uvarenko/todo-list/internal/controller/http/rest"
	"github.com/uvarenko/todo-list/internal/repo"
	"github.com/uvarenko/todo-list/internal/usecase"
	"github.com/uvarenko/todo-list/pkg/httpserver"
	"github.com/uvarenko/todo-list/pkg/logger"
	"github.com/uvarenko/todo-list/pkg/postgre"
)

func Run(l *logger.Logger) {
	// init env
	godotenv.Load()

	// db
	db := postgre.New(l)

	// use case
	todoUseCase := usecase.NewToDoUseCase(
		repo.New(db),
	)

	// http server
	handler := gin.New()
	rest.NewRouter(handler, todoUseCase)
	httpServer := httpserver.New(handler)

	// shutdown logic
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app was interrupted with singnal = " + s.String())
	case err := <-httpServer.Start():
		l.Error("app error = %w", err)
	}

	err := httpServer.ShutDown()
	if err != nil {
		l.Error("shutdown with error = %w", err)
	}
}
