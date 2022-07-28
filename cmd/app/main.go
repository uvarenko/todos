package main

import (
	"github.com/spf13/cobra"
	"github.com/uvarenko/todo-list/cmd"
	"github.com/uvarenko/todo-list/internal/app"
	"github.com/uvarenko/todo-list/pkg/logger"
)

func main() {
	l := logger.New("debug")
	err := cmd.Execute(func(cmd *cobra.Command, args []string) {
		app.Run(l)
	})
	if err != nil {
		l.Fatal("could not start the app")
	}
}
