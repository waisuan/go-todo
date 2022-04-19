package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Router", func() {
	var router *gin.Engine
	var todoList *TodoList

	BeforeEach(func() {
		todoList = NewTodoList()
		_, err := todoList.InsertTodoItem("Title", "Description", "2020-09-12 19:12")
		Expect(err).To(BeNil())

		router = SetupRouter(todoList)
	})

	Context("Get todo list items", func() {
		It("should return a list of items", func() {
			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/todo", nil)
			Expect(err).To(BeNil())

			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))

			var response []TodoItem
			err = json.NewDecoder(w.Body).Decode(&response)
			Expect(err).To(BeNil())
			Expect(len(response)).To(Equal(1))
			Expect(response[0].String()).To(Equal("Title | Description | 2020-09-12 19:12:00 +0000 UTC"))
		})
	})

	Context("Insert todo item", func() {
		It("should insert a new item successfully", func() {
			w := httptest.NewRecorder()
			reqBody := Request{
				Title:       "Hello",
				Description: "World",
				DueDateTime: "2020-01-08 13:09",
			}
			reqBodyAsJson, _ := json.Marshal(reqBody)
			req, err := http.NewRequest("POST", "/todo/item", bytes.NewBuffer(reqBodyAsJson))
			Expect(err).To(BeNil())

			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))

			var response *TodoItem
			err = json.NewDecoder(w.Body).Decode(&response)
			Expect(err).To(BeNil())
			Expect(response.Id).ToNot(BeNil())
			Expect(response.Title).To(Equal("Hello"))
			Expect(response.Description).To(Equal("World"))
			Expect(response.DueDateTime.String()).To(Equal("2020-01-08 13:09:00 +0000 UTC"))
			Expect(response.CreatedAt).ToNot(BeNil())
		})

		Context("When given an invalid request body", func() {
			It("should return a 4XX error code", func() {
				w := httptest.NewRecorder()
				reqBody := Request{
					Title:       "Hello",
					Description: "World",
					DueDateTime: "2020-01-08T13:09",
				}
				reqBodyAsJson, _ := json.Marshal(reqBody)
				req, err := http.NewRequest("POST", "/todo/item", bytes.NewBuffer(reqBodyAsJson))
				Expect(err).To(BeNil())

				router.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(ContainSubstring("invalid due date time given"))
			})
		})
	})
})
