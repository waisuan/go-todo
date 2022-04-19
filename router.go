package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDateTime string `json:"dueDateTime"`
}

type RouteHandler struct {
	todoList *TodoList
}

func SetupRouter(todoList *TodoList) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	handler := RouteHandler{todoList: todoList}
	r.GET("/todo", handler.getTodoList)
	r.POST("/todo/item", handler.insertTodoItem)

	return r
}

func (h *RouteHandler) getTodoList(c *gin.Context) {
	handleValidRequest(c, h.todoList.GetAllTodos())
}

func (h *RouteHandler) insertTodoItem(c *gin.Context) {
	var request *Request
	err := c.BindJSON(&request)
	if err != nil {
		handleRequestError(c, err)
		return
	}

	newItem, err := h.todoList.InsertTodoItem(request.Title, request.Description, request.DueDateTime)
	if err != nil {
		handleRequestError(c, err)
		return
	}

	handleValidRequest(c, newItem)
}

func handleRequestError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, err.Error())
}

func handleValidRequest(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}
