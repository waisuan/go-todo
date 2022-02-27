package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"io"
)

type Prompt interface {
	Run(label string) (string, error)
	RunOnSelect(label string, items []TodoItem) (int, string, error)
}

type PromptWrapper struct{}

func (pw *PromptWrapper) Run(label string) (string, error) {
	prompt := promptui.Prompt{Label: label}
	return prompt.Run()
}

func (pw *PromptWrapper) RunOnSelect(label string, items []TodoItem) (int, string, error) {
	prompt := promptui.Select{Label: label, Items: items}
	return prompt.Run()
}

func PromptToAddItem(todoList *TodoList, prompt Prompt) error {
	title, err := prompt.Run("Title")
	if err != nil {
		return err
	}

	description, err := prompt.Run("Description")
	if err != nil {
		return err
	}

	dueDateTime, err := prompt.Run("Due Date Time (YYYY-MM-DD HH:MM)")
	if err != nil {
		return err
	}

	_, err = todoList.InsertTodoItem(title, description, dueDateTime)
	if err != nil {
		return err
	}

	return nil
}

func PromptToListAllItems(todoList *TodoList, writer io.Writer) error {
	for _, v := range todoList.GetAllTodos() {
		_, err := fmt.Fprintf(writer, "Title=%s, Desc=%s, DueDateTime=%s\n", v.title, v.description, v.dueDateTime)
		if err != nil {
			return err
		}
	}

	return nil
}

func PromptToDeleteItem(todoList *TodoList, prompt Prompt) error {
	todos := todoList.GetAllTodos()
	idx, _, err := prompt.RunOnSelect("Which item do you want to delete?", todos)
	if err != nil {
		return err
	}

	err = todoList.DeleteTodoItem(todos[idx].id)
	if err != nil {
		return err
	}

	return nil
}
