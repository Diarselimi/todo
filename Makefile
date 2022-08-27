
build:
	go build cmd/main.go

help:
	echo "List of actions:"
	echo "add 'Todo text here'; adds a new todo in the list"
	echo "list; Shows all the todos you have"
	echo "complete id;Completes one todo with the given id"
	echo "clean; Cleans up the completed todos"

