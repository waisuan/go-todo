package main

import (
	"log"
	"net/http"
)

func main() {
	todoList := NewTodoList()
	err := http.ListenAndServe(":8080", NewRouteHandler(todoList))
	if err != nil {
		log.Fatalln(err)
	}
}
