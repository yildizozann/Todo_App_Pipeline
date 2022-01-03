package todo_test

import (
	"testing"
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"time"
	"todo_app/config"
	"todo_app/internal/todo"

	"github.com/stretchr/testify/assert"
)

func TestAddNewToDoToDatabese(t *testing.T){
	config.ConnectDB()
	requestStruct := todo.AddToDoRequest{
		Todo: "go to swim",
	}
	database := todo.NewDatabase()
	result, err := database.AddNewToDoToDatabase(&requestStruct)

	var expectedResult todo.AddToDoResponse
	expectedResult.Todo = requestStruct.Todo
	expectedResult.Message = "New ToDo Added to Databese"
	
	assert.Equal(t, nil, err )
	assert.Equal(t, expectedResult, *result)

	
}
func TestGetAllToDo(t *testing.T){
	config.ConnectDB()
	collection := config.MI.DB.Collection("todo")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r := todo.AddToDoRequest{
		Todo: "finish the task",
	}
	database := todo.NewDatabase()
	count, _ := collection.CountDocuments(ctx, bson.D{})
	_, err := collection.InsertOne(ctx, r)

	assert.Nil(t, err)
	result, _ := database.GetAllTodo()
	expectedCount := count + 1

	assert.Equal(t, expectedCount, int64(len(*result)))
} 