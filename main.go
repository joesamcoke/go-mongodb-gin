package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoById)
	router.POST("/todos", createTodo)
	router.DELETE("/todos/:id", deleteTodo)
	router.PUT("/todos/:id", updateTodo)

	router.Run("localhost:8080")
}
