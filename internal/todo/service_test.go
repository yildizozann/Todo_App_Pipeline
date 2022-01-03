package todo_test

import (
	"context"
	"testing"
	mocks "todo_app/internal/mocks/todo"
	"todo_app/internal/todo"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

/* func TestAddNewToDo(t *testing.T){
  config.ConnectDB()

  request := todo.AddToDoRequest{
	  Todo: "cook dinner",
  }
  expextedResponse := todo.AddToDoResponse{
    Todo: "cook dinner",
    Message: "New ToDo Added to Databese",
  }
  database := todo.NewDatabase()
  service := todo.NewService(database)

  response, _ := service.AddNewToDo(context.Background(),&request)


  assert.Equal(t, expextedResponse, *response )

} */
func TestAddNewToDo(t *testing.T){
  ctrl:= gomock.NewController(t)
  defer ctrl.Finish()
  request := todo.AddToDoRequest{
	  Todo: "cook dinner",
  }
  expextedResponse := todo.AddToDoResponse{
    Todo: "cook dinner",
    Message: "New ToDo Added to Databese",
  }
  mockDatabase:= mocks.NewMockServiceDatabase(ctrl)
  mockDatabase.EXPECT().AddNewToDoToDatabase(&request).Return(&expextedResponse,nil)
  service:= todo.NewService(mockDatabase)

  response, _ := service.AddNewToDo(context.Background(),&request)

  assert.Equal(t, expextedResponse, *response )

}

func TestGetAll(t *testing.T){
  ctrl:= gomock.NewController(t)
  defer ctrl.Finish()
  list := []todo.DatabaseStruct{}
	list = append(list, 
		todo.DatabaseStruct{
			ID:   [12]byte{},
			Todo: "go to shop",
		},
		todo.DatabaseStruct{
			ID:   [12]byte{},
			Todo: "dummy text!",
		})


  expectedResponse := list
  mockDatabase:= mocks.NewMockServiceDatabase(ctrl)
  mockDatabase.EXPECT().GetAllTodo().Return(&expectedResponse,nil)
  service:= todo.NewService(mockDatabase)
  
  response := service.GetAll()

  expextedResponseFromService := todo.GetAllResponse{
    Todolist: expectedResponse,
    Message: "Success",
  }

  assert.Equal(t, expextedResponseFromService, *response)
}