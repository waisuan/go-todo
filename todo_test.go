package main

import (
	"reflect"
	"testing"
	"time"
)

func TestTodoList(t *testing.T) {
	createTodoItem := func(todoList *TodoList) *TodoItem {
		t.Helper()

		todoItem, err := todoList.InsertTodoItem("title", "description", "2006-01-02 15:04")
		if err != nil {
			t.Fatalf("unable to insert todo item: %v", err)
		}

		return todoItem
	}

	t.Run("get empty todo list", func(t *testing.T) {
		todoList := NewTodoList()

		var want []TodoItem
		got := todoList.GetAllTodos()

		if len(want) != len(got) {
			t.Errorf("expected an empty todo list")
		}
	})

	t.Run("new todo items have IDs", func(t *testing.T) {
		todoList := NewTodoList()

		got := createTodoItem(todoList).Id

		if got == "" {
			t.Errorf("expected a non-empty item ID")
		}
	})

	t.Run("get todo items in ascending order", func(t *testing.T) {
		todoList := NewTodoList()

		var want []TodoItem
		for i := 1; i <= 5; i++ {
			todoItem := *createTodoItem(todoList)
			want = append(want, todoItem)
		}

		got := todoList.GetAllTodos()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected some todo items: want: %v, got: %v", want, got)
		}
	})

	t.Run("insert with invalid due date time", func(t *testing.T) {
		todoList := NewTodoList()

		_, err := todoList.InsertTodoItem("title", "description", "2006-01-02")
		if err == nil {
			t.Errorf("expected an error with the given due date time")
		}
	})

	t.Run("get specific todo item", func(t *testing.T) {
		todoList := NewTodoList()

		id := createTodoItem(todoList).Id

		_, err := todoList.GetTodoItem(id)
		if err != nil {
			t.Errorf("expected a todo item that exists")
		}
	})

	t.Run("get non-existent todo item", func(t *testing.T) {
		todoList := NewTodoList()

		_, err := todoList.GetTodoItem("id")
		if err == nil {
			t.Errorf("did not expect to find todo item")
		}
	})

	t.Run("delete a todo item", func(t *testing.T) {
		todoList := NewTodoList()

		id := createTodoItem(todoList).Id

		err := todoList.DeleteTodoItem(id)
		if err != nil {
			t.Fatalf("expected a todo item that exists")
		}

		_, err = todoList.GetTodoItem(id)
		if err == nil {
			t.Errorf("did not expect to find todo item")
		}
	})

	t.Run("delete a non-existent todo item", func(t *testing.T) {
		todoList := NewTodoList()

		err := todoList.DeleteTodoItem("id")
		if err == nil {
			t.Errorf("expected to not be able to delete todo item")
		}
	})

	t.Run("update a todo item", func(t *testing.T) {
		todoList := NewTodoList()

		todoItem := *createTodoItem(todoList)

		todoItem.Description = "Updated Description"
		err := todoList.UpdateTodoItem(todoItem)
		if err != nil {
			t.Fatalf("expected a todo item that exists")
		}

		item, _ := todoList.GetTodoItem(todoItem.Id)
		if item.Description != todoItem.Description {
			t.Errorf("expected todo item to be updated")
		}
	})

	t.Run("update a non-existent todo item", func(t *testing.T) {
		todoList := NewTodoList()

		err := todoList.UpdateTodoItem(TodoItem{})
		if err == nil {
			t.Errorf("expected to not be able to update todo item")
		}
	})

	t.Run("get due todo items", func(t *testing.T) {
		todoList := NewTodoList()

		yesterday := time.Now().Add(-(time.Hour * 24))
		nextDay := time.Now().Add(time.Hour * 24)
		_, _ = todoList.InsertTodoItem("title", "description", yesterday.Format("2006-01-02 15:04"))
		item2, _ := todoList.InsertTodoItem("title", "description", nextDay.Format("2006-01-02 15:04"))

		want := []TodoItem{*item2}
		got := todoList.GetDueTodoItems()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected due todo items: got %v, want %v", got, want)
		}
	})

	t.Run("no due todo items", func(t *testing.T) {
		todoList := NewTodoList()

		yesterday := time.Now().Add(-(time.Hour * 24))
		_, _ = todoList.InsertTodoItem("title", "description", yesterday.Format("2006-01-02 15:04"))

		got := todoList.GetDueTodoItems()
		if len(got) != 0 {
			t.Errorf("expected to not have any due todo items")
		}
	})
}
