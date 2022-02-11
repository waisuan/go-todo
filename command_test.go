package main

import (
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
