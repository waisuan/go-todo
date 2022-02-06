package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "How can I help?",
		Items: []string{"Add todo item", "List all todo items"},
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Something went wrong: %v\n", err)
		return
	}

	todoList := NewTodoList()
	if i == 1 {
		fmt.Println("Searching for todo items...")
		for v := range todoList.GetAllTodos() {
			fmt.Println(v)
		}
	} else {
		fmt.Println("Operation is not supported yet!")
	}
}
