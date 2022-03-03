package main

import (
	"fmt"
	"testing"

	"github.com/joesamcoke/go-mongodb/models"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var testTodoID string

// Run tests all at once

func TestCreateTodo(t *testing.T) {

	Todo := models.Todo{Title: "this is a test todo", Completed: true}

	result, err := service.CreateTodo(Todo)
	if err != nil {
		t.Errorf("Expected Todo, received %v", err)
	}

	testTodoID = result.InsertedID.(primitive.ObjectID).Hex()

	fmt.Println(testTodoID)

	require.Nil(t, err, "No error - Pass")
	require.NotNil(t, result, "We have a result - Pass")
	require.NotNil(t, testTodoID, "We an inserted ID - Pass")
}

func TestGetTodos(t *testing.T) {
	result, err := service.GetTodos()
	if err != nil {
		t.Errorf("Expected Todos, received %v", err)
	}

	require.Nil(t, err, "No error - Pass")
	require.NotNil(t, result, "We have a result - Pass")
	require.Equal(t, result[len(result)-1].ID, testTodoID, "ID is correct - Pass")
}

func TestGetTodoById(t *testing.T) {
	result, err := service.GetTodoById(testTodoID)
	if err != nil {
		t.Errorf("Expected Todos, received %v", err)
	}

	require.Nil(t, err, "No errors - Pass")
	require.NotNil(t, result, "We have a result - Pass")
	require.Equal(t, result.ID, testTodoID, "ID is correct - Pass")
}

func TestUpdateTodo(t *testing.T) {

	Todo := models.Todo{Title: "this is an updated test todo", Completed: false}

	result, err := service.UpdateTodo(testTodoID, Todo)
	if err != nil {
		t.Errorf("Expected Todo, received %v", err)
	}

	var expectedCount int64 = 1

	require.Nil(t, err, "No error - Pass")
	require.NotNil(t, result, "We have a result - Pass")
	require.Equal(t, result.ModifiedCount, expectedCount, "Update is correct - Pass")
}

func TestDeleteTodo(t *testing.T) {

	result, err := service.DeleteTodo(testTodoID)
	if err != nil {
		t.Errorf("Expected Todo, received %v", err)
	}

	var expectedCount int64 = 1

	require.Nil(t, err, "No error - Pass")
	require.NotNil(t, result, "We have a result - Pass")
	require.Equal(t, result.DeletedCount, expectedCount, "Delete is correct - Pass")
}
