package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

func main() {
	todoList := NewTodoList()
	for {
		prompt := promptui.Select{
			Label: "How can I help?",
			Items: []string{"Add todo item", "List all todo items", "Delete a todo item", "Exit"},
		}

		i, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Something went wrong: %v\n", err)
			break
		}

		if i == 0 {
			err = PromptToAddItem(todoList, &PromptWrapper{})
		} else if i == 1 {
			err = PromptToListAllItems(todoList, os.Stdout)
		} else if i == 2 {
			err = PromptToDeleteItem(todoList, &PromptWrapper{})
		} else if i == 3 {
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
