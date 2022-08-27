package command

import (
	"cards/internal/infrastructure"
	"fmt"
)

type CleanCompletedTodosCommand struct {
	storage infrastructure.IStorage
}

func NewCleanCompletedTodosCommand() CleanCompletedTodosCommand {
	return CleanCompletedTodosCommand{
		storage: infrastructure.NewStorage(),
	}
}

func (c CleanCompletedTodosCommand) Exec() {
	todos, err := c.storage.GetList()
	if err != nil {
		fmt.Println("Data could not be retrived.")
	}

	todos = todos.CleanCompleted()

	c.storage.Save(todos)
}
