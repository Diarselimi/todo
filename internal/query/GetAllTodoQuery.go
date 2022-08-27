package query

import (
	"cards/internal/domain"
	"cards/internal/infrastructure"
	"fmt"
)

type Todo struct {
	id     int
	value  string
	isDone bool
}

type GetAllTodo []Todo

type GetAllTodoQuery struct {
	storage infrastructure.IStorage
	filters string
}

func NewGetAllTodoQuery(filters string) GetAllTodoQuery {
	return GetAllTodoQuery{
		storage: infrastructure.NewStorage(),
		filters: filters,
	}
}

func (t GetAllTodoQuery) Get() domain.TodoSlice {
	list, err := t.storage.GetList()
	if err != nil {
		fmt.Println("Could not get data")
		panic("Aaaaaa")
	}

	list.Print()

	return list
}
