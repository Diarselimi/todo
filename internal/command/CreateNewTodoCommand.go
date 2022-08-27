package command

import (
	"cards/internal/infrastructure"
	"fmt"
)

type CreateNewTodo struct {
	title       string
	description string
}

type CreateNewTodoCommand struct {
	storage infrastructure.IStorage
	command CreateNewTodo
}

func NewCreateNewTodoCommand(title string) CreateNewTodoCommand {
	return CreateNewTodoCommand{
		storage: infrastructure.NewStorage(),
		command: CreateNewTodo{title, title},
	}
}

func (c CreateNewTodoCommand) Exec() {
	todos, err := c.storage.GetList()
	fmt.Println(todos)
	if err != nil {
		fmt.Println("Could not get data from storage")
	}

	todos = todos.Add(c.command.title)
	c.storage.Save(todos)

	fmt.Println("Create New TODO Executed")
}
