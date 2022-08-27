package config

import (
	"cards/internal/command"
	"cards/internal/query"
	"fmt"
	"os"
)

type Route map[string]func(data string)

func addNewTodo(data string) {
	command.NewCreateNewTodoCommand(data).Exec()
	fmt.Println("Done!")
}

func deleteTodo(data string) {
	command.NewCompleteTodoCommand(data).Exec()
}

func cleanCompletedTodos(data string) {
	command.NewCleanCompletedTodosCommand().Exec()
}

func getAllTodo(data string) {
	query.NewGetAllTodoQuery(data).Get()
}

func NewRouter() Route {
	return Route{
		"list":     getAllTodo,
		"add":      addNewTodo,
		"complete": deleteTodo,
		"clean":    cleanCompletedTodos,
	}
}

func (r Route) GetHandler(route string) func(data string) {
	fn, ok := r[route]
	if !ok {
		fmt.Println("Route not found!")
		os.Exit(1)
	}
	return fn
}
