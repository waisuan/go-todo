package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouteHandler(t *testing.T) {
	todoList := NewTodoList()
	handler := NewRouteHandler(todoList)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	createTodoItem := func(todoList *TodoList) *TodoItem {
		t.Helper()

		todoItem, err := todoList.InsertTodoItem("title", "description", "2006-01-02 15:04")
		if err != nil {
			t.Fatalf("unable to insert todo item: %v", err)
		}

		return todoItem
	}

	t.Run("Get todo list request", func(t *testing.T) {
		todoItem := createTodoItem(todoList)

		res, err := http.Get(ts.URL + "/todos")
		if err != nil {
			t.Fatalf("expected no errors but got %v", err)
		}

		var result []TodoItem
		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			t.Fatalf("expected no errors but got %v", err)
		}

		if len(result) != 1 && result[0] != *todoItem {
			t.Errorf("expected %v but got %v", todoItem, result)
		}
	})

	t.Run("Insert todo item request", func(t *testing.T) {
		reqBody := Request{
			Title:       "title",
			Description: "description",
			DueDateTime: "2006-01-02 15:04",
		}
		b, _ := json.Marshal(reqBody)
		res, _ := http.Post(ts.URL+"/todos", "application/json", bytes.NewBuffer(b))

		var todoItem TodoItem
		_ = json.NewDecoder(res.Body).Decode(&todoItem)
		fmt.Println(todoItem)
	})
}
