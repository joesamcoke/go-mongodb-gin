package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/joesamcoke/go-mongodb/models"
	"github.com/joesamcoke/go-mongodb/services"
)

var service = services.NewTodoService()

// Get all Todos
func getTodos(c *gin.Context) {
	service, err := service.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Format for JSON response
	jsonData, err := json.Marshal(service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Return Data
	c.Data(http.StatusOK, "application/json", jsonData)
}

// Get Todo by ID
func getTodoById(c *gin.Context) {
	id := c.Param("id")

	todo, err := service.GetTodoById(id)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot retieve result"})
	}
	fmt.Println(todo)
	// Format for JSON response
	jsonData, err := json.Marshal(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Return Data
	c.Data(http.StatusOK, "application/json", jsonData)
}

// Create new todo
func createTodo(c *gin.Context) {
	Todo := models.NewTodo()

	// Decode input to JSON
	err := json.NewDecoder(c.Request.Body).Decode(&Todo)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect input"})
		return
	}

	// Validate against struct
	var validate = validator.New()
	if err := validate.Struct(Todo); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect input"})
		return
	}

	// Insert todo
	result, err := service.CreateTodo(Todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		log.Panic(err)
	}

	// Format for JSON response
	jsonData, err := json.Marshal(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		log.Panic(err)
	}

	// Return inserted ID
	c.Data(http.StatusOK, "application/json", jsonData)

}

// Delete todo
func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	result, err := service.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete item"})
		log.Panic(err)
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete item"})
		log.Panic(err)
	}

	c.Data(http.StatusOK, "application/json", jsonData)

}

func updateTodo(c *gin.Context) {
	Todo := models.NewTodo()
	id := c.Param("id")

	err := json.NewDecoder(c.Request.Body).Decode(&Todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect input"})
		return
	}

	// Validate against struct
	var validate = validator.New()
	if err := validate.Struct(&Todo); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect input"})
		return
	}

	fmt.Println(Todo)

	// Update todo
	result, err := service.UpdateTodo(id, Todo)
	if err != nil {
		fmt.Println(err)
	}

	// Format for JSON response
	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Panic(err)
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}
