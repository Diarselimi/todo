package command

import (
	"cards/internal/infrastructure"
	"fmt"
	"strconv"
)

type CompleteTodoCommand struct {
	storage infrastructure.IStorage
	todoKey int
}

func NewCompleteTodoCommand(val string) CompleteTodoCommand {
	number, err := strconv.Atoi(val)
	if err != nil {
		panic("Input is not valid")
	}
	return CompleteTodoCommand{storage: infrastructure.NewStorage(), todoKey: number}
}

func (c CompleteTodoCommand) Exec() {
	//some steps
	todos, err := c.storage.GetList()
	if err != nil {
		fmt.Println("There was an error while fetching todos.")
		panic(err)
	}
	todos.Print()
	todos = todos.Complete(c.todoKey)
	c.storage.Save(todos)
	fmt.Printf("Todo %d ✅ Completed", c.todoKey)
}
