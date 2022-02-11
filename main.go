package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func main() {
	todoList := NewTodoList()
	for {
		prompt := promptui.Select{
			Label: "How can I help?",
			Items: []string{"Add todo item", "List all todo items", "Exit"},
		}

		i, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Something went wrong: %v\n", err)
			break
		}

		if i == 0 {
			err = PromptToAddItem(todoList, &PromptWrapper{})
		} else if i == 1 {
			printAllTodoItems(todoList)
		} else if i == 2 {
			break
		} else {
			fmt.Println("Operation is not supported yet!")
		}

		if err != nil {
			fmt.Printf("Something went wrong: %v\n", err)
			break
		}
	}
}

func printAllTodoItems(todoList *TodoList) {
	fmt.Println("Searching for todo items...")
	for _, v := range todoList.GetAllTodos() {
		fmt.Printf("Title=%s, Desc=%s, DueDateTime=%s\n", v.title, v.description, v.dueDateTime)
	}
}
