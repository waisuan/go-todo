package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestPromptToAddItem(t *testing.T) {
	t.Run("the right prompts were triggered", func(t *testing.T) {
		todoList := NewTodoList()

		mockPromptWrapper := MockPromptWrapper{}
		got := PromptToAddItem(todoList, &mockPromptWrapper)
		if got != nil {
			t.Fatalf("did not expect to see an error")
		}

		labels := mockPromptWrapper.labels
		want := []string{"Title", "Description", "Due Date Time (YYYY-MM-DD HH:MM)"}
		if !reflect.DeepEqual(labels, want) {
			t.Errorf("expected the rigth set of prompts. want: %v, got: %v", want, labels)
		}
	})

	t.Run("expected to insert a todo item", func(t *testing.T) {
		todoList := NewTodoList()

		got := PromptToAddItem(todoList, &MockPromptWrapper{})
		if got != nil {
			t.Fatalf("did not expect to see an error")
		}

		if len(todoList.GetAllTodos()) != 1 {
			t.Errorf("expected just one todo item to be inserted")
		}
	})
}

func TestPromptToListAllItems(t *testing.T) {
	t.Run("prints all todo items", func(t *testing.T) {
		todoList := NewTodoList()
		_, _ = todoList.InsertTodoItem("Title", "Description", "1991-11-12 13:43")
		_, _ = todoList.InsertTodoItem("Title2", "Description2", "1992-11-12 13:43")

		buffer := bytes.Buffer{}
		err := PromptToListAllItems(todoList, &buffer)
		if err != nil {
			t.Fatalf("did not expect to see an error")
		}

		want := `Title=Title, Desc=Description, DueDateTime=1991-11-12 13:43:00 +0000 UTC
Title=Title2, Desc=Description2, DueDateTime=1992-11-12 13:43:00 +0000 UTC
`
		got := buffer.String()
		if got != want {
			t.Errorf("expected all todo items to be printed. want: %v, got: %v", want, got)
		}
	})
}
