package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Router", func() {
	var router *gin.Engine

	BeforeEach(func() {
		router = SetupRouter()
	})

	Context("Get todo list items", func() {
		It("should run", func() {
			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/todos", nil)
			Expect(err).To(BeNil())

			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})
})
