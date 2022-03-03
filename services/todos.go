package services

import (
	"context"
	"fmt"
	"time"

	"github.com/joesamcoke/go-mongodb/db"
	"github.com/joesamcoke/go-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoService struct {
	Coll   *mongo.Collection
	Client *mongo.Client
}

func NewTodoService() TodoService {
	client, db := db.MongoCon()
	collection := "todos"
	return TodoService{
		Coll:   client.Database(db).Collection(collection),
		Client: client,
	}

}

// Get all Todos
func (service TodoService) GetTodos() ([]models.Todo, error) {

	var Todos []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := service.Coll.Find(ctx, bson.D{})

	if err != nil {
		return Todos, err
	}

	// Loop & decode results
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var todo models.Todo
		if err = cursor.Decode(&todo); err != nil {
			return Todos, err
		}
		Todos = append(Todos, todo)
	}

	return Todos, err
}

// Get Todo by ID
func (service TodoService) GetTodoById(id string) (models.Todo, error) {

	Todo := models.NewTodo()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Convert id string to bson object
	oId, _ := primitive.ObjectIDFromHex(id)

	// Get result & decode
	err := service.Coll.FindOne(ctx, bson.M{"_id": oId}).Decode(&Todo)
	if err != nil {
		return Todo, err
	}

	return Todo, err
}

// Create new todo
func (service TodoService) CreateTodo(Todo models.Todo) (*mongo.InsertOneResult, error) {
	// Todo := models.NewTodo()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println(Todo)

	result, err := service.Coll.InsertOne(ctx, Todo)
	if err != nil {
		fmt.Println(err)
	}

	// Return the ID of the result if successful
	return result, err

}

func (service TodoService) DeleteTodo(id string) (*mongo.DeleteResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
	}

	result, err := service.Coll.DeleteOne(ctx, bson.M{"_id": oId})

	return result, err
}

func (service TodoService) UpdateTodo(id string, Todo models.Todo) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oId, _ := primitive.ObjectIDFromHex(id)

	update := bson.M{"title": Todo.Title, "completed": Todo.Completed}

	result, err := service.Coll.UpdateOne(ctx, bson.M{"_id": oId}, bson.M{"$set": update})

	return result, err
}
