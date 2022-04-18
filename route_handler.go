package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDateTime string `json:"dueDateTime"`
}

func NewRouteHandler(todoList *TodoList) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/todos", func(writer http.ResponseWriter, request *http.Request) {
		getTodoList(writer, request, todoList)
	}).Methods("GET")

	router.HandleFunc("/todos", func(writer http.ResponseWriter, request *http.Request) {
		insertTodoItem(writer, request, todoList)
	}).Methods("POST")

	return router
}

func getTodoList(w http.ResponseWriter, r *http.Request, todoList *TodoList) {
	err := json.NewEncoder(w).Encode(todoList.GetAllTodos())
	if err != nil {
		http.Error(w, "Something went wrong!", 400)
	}
}

func insertTodoItem(w http.ResponseWriter, r *http.Request, todoList *TodoList) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var req Request
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		http.Error(w, "Something went wrong!", 400)
	}

	item, err := todoList.InsertTodoItem(req.Title, req.Description, req.DueDateTime)
	if err != nil {
		http.Error(w, "Something went wrong!", 400)
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, "Something went wrong!", 400)
	}
}
