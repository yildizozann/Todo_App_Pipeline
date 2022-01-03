package todo

import (
	"context"
	"time"
	"todo_app/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type DatabaseStruct struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Todo string `json:"todo,omitempty" bson:"todo,omitempty"`
} 

const Collection = "todo"

type Database struct {

}

func NewDatabase() *Database {

	database := Database{}

	return &database
}

func (d *Database) AddNewToDoToDatabase(r *AddToDoRequest) (*AddToDoResponse, error ){

	collection := config.MI.DB.Collection(Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, r)
	if err != nil {
		return nil,err
	}else {
		return &AddToDoResponse{
			Todo: r.Todo,
			Message: "New ToDo Added to Databese",
		}, nil
	}
} 
func (d *Database) GetAllTodo() (*[]DatabaseStruct, error) {
	collection := config.MI.DB.Collection(Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var response []DatabaseStruct

	cursor, err := collection.Find(ctx, bson.D{})
	

	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx){
		var todo DatabaseStruct
		cursor.Decode(&todo)
		response = append(response, todo)
	}
	return &response, nil
}

