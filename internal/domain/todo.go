package domain

import (
	"fmt"
)

type Todo struct {
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

type TodoSlice []Todo

func (todos TodoSlice) Print() {

	for _, todo := range todos {
		is := ""
		if todo.Completed {
			is = "✅"
		}
		fmt.Printf("[%d] %s %s \n\n", todo.Id, todo.Title, is)
	}
}

func NewTodoFromArray(todos []string) TodoSlice {
	if len(todos) <= 0 {
		return TodoSlice{}
	}

	var todoCollection TodoSlice
	for i, todo := range todos {
		todoCollection = append(todoCollection, Todo{i, todo, false})
	}
	return todoCollection
}

func (t TodoSlice) First() Todo {
	return t[0]
}

func (t TodoSlice) Add(val string) TodoSlice {
	return append(t, Todo{len(t), val, false})
}

func (t TodoSlice) CleanCompleted() TodoSlice {
	todos := TodoSlice{}
	for _, todo := range t {
		if todo.Completed == false {
			todos = append(todos, todo)
		}
	}
	return todos
}

func (t TodoSlice) Complete(id int) TodoSlice {
	for i, todo := range t {
		if todo.Id == id {
			t[i].Completed = true
		}
	}
	return t
}
