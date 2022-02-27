package main

import "strings"

type MockPromptWrapper struct {
	labels []string
	items  *[]TodoItem
}

func (pw *MockPromptWrapper) Run(label string) (string, error) {
	pw.labels = append(pw.labels, label)

	if strings.Contains(label, "YYYY-MM-DD HH:MM") {
		return "1991-11-12 11:30", nil
	}

	return label, nil
}

func (pw *MockPromptWrapper) RunOnSelect(label string, items []TodoItem) (int, string, error) {
	pw.labels = append(pw.labels, label)
	pw.items = &items

	return 0, "some_item", nil
}
