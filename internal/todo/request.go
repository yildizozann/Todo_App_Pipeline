package todo

import "go.mongodb.org/mongo-driver/bson/primitive"

type AddToDoRequest struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Todo string `json:"todo,omitempty" bson:"todo,omitempty"`
}

