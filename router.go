package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/todos", getTodoList)

	return r
}

func getTodoList(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"msg": "Yo!"})
}
