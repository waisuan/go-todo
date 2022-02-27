package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"sort"
	"time"
)

type TodoItem struct {
	id          string
	title       string
	description string
	dueDateTime time.Time
	createdAt   time.Time
}

type TodoList struct {
	items map[string]TodoItem
}

const dueDateLayout = "2006-01-02 15:04"

func NewTodoList() *TodoList {
	return &TodoList{
		items: make(map[string]TodoItem),
	}
}

func (t *TodoList) GetAllTodos() []TodoItem {
	var items []TodoItem
	for k := range t.items {
		items = append(items, t.items[k])
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].createdAt.Before(items[j].createdAt)
	})
	return items
}

func (t *TodoList) GetTodoItem(id string) (TodoItem, error) {
	item, ok := t.items[id]
	if !ok {
		return item, errors.New("item does not exist")
	}

	return item, nil
}

func (t *TodoList) InsertTodoItem(title string, description string, dueDateTime string) (*TodoItem, error) {
	ti, err := time.Parse(dueDateLayout, dueDateTime)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid due date time given: %v", err))
	}

	item := TodoItem{
		id:          uuid.New().String(),
		title:       title,
		description: description,
		dueDateTime: ti,
		createdAt:   time.Now(),
	}

	t.items[item.id] = item

	return &item, nil
}

func (t *TodoList) UpdateTodoItem(item TodoItem) error {
	_, ok := t.items[item.id]
	if !ok {
		return errors.New("can't update non-existent item")
	}

	t.items[item.id] = item

	return nil
}

func (t *TodoList) DeleteTodoItem(id string) error {
	_, ok := t.items[id]
	if !ok {
		return errors.New("can't delete non-existent item")
	}

	delete(t.items, id)

	return nil
}

func (t *TodoList) GetDueTodoItems() []TodoItem {
	var dueItems []TodoItem
	now := time.Now()
	for _, item := range t.items {
		if now.Before(item.dueDateTime) {
			dueItems = append(dueItems, item)
		}
	}

	return dueItems
}

func (tt TodoItem) String() string {
	return fmt.Sprintf("%v | %v | %v", tt.title, tt.description, tt.dueDateTime)
}
