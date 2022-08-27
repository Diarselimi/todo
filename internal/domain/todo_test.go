package domain

import (
	"testing"
)

func TestTodo_newTodoFromMissingFile(t *testing.T) {
	todo := NewTodoFromArray([]string{})
	if len(todo) != 0 {
		t.Errorf("Expected no items in a list of todos, but got %d ", len(todo))
	}
}

func TestSavingTodosIntoNewFile(t *testing.T) {

	todo := NewTodoFromArray([]string{"A", "B", "C"})

	if len(todo) != 3 {
		t.Errorf("Expected number of todos %d, but got %d", 3, len(todo))
	}
}

func TestTodo_First(t *testing.T) {
	todo := NewTodoFromArray([]string{"A", "B"})
	if todo.First() != "A" {
		t.Errorf("Expected value of %s, but got %s", "A", todo.First())
	}
}

func TestTodo_ShuffleTodoItems(t *testing.T) {
	todo := NewTodoFromArray([]string{"A", "B", "C"})
	newTodo := NewTodoFromArray([]string{"A", "B", "C"})
	todo.shuffle()

	if todo.toString() == newTodo.toString() {
		t.Errorf("Expected items to be random %s, but got same %s", todo.toString(), newTodo.toString())
	}
}
