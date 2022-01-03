package todo

import (
	"context"
)

type ServiceDatabase interface {
	AddNewToDoToDatabase(r *AddToDoRequest) (*AddToDoResponse, error )
	GetAllTodo() (*[]DatabaseStruct, error)	
}

type Service struct {
	Database ServiceDatabase
}

func NewService(d ServiceDatabase) *Service {

	service := Service{
		Database: d,
	}
	return &service
} 

func (s *Service) AddNewToDo(ctx context.Context, request *AddToDoRequest) (*AddToDoResponse, error){
	result,err := s.Database.AddNewToDoToDatabase(request)

	if err == nil {
		return result , nil
	}else {
		return &AddToDoResponse{}, err
	}
}
func (s *Service) GetAll() *GetAllResponse{
	result, err := s.Database.GetAllTodo()
	if err != nil {
		return &GetAllResponse{
			Message: "Do Not Have Any To Do",
		}
	}else{
		return &GetAllResponse{
			Todolist: *result,
			Message: "Success",
		}
	}
}
