package main

import (
	"cards/internal/config"
	"fmt"
	"os"
)

func main() {
	router := config.NewRouter()

	var route, value string
	route = os.Args[1]
	value = os.Args[2]

	if value == "" {
		fmt.Println("\n Please add the value after command `/command value`")
		os.Exit(1)
	}
	handler := router.GetHandler(route)
	handler(value)
}
