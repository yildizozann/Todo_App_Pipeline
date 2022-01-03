package todo_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	mocks "todo_app/internal/mocks/todo"
	"todo_app/internal/todo"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

  func TestHealth(t *testing.T){
	app := todo.CreateApp()
	database := todo.NewDatabase()
	service := todo.NewService(database)
	handler := todo.NewHandler(service)

	handler.RegisterRoutes(app)

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	request.Header.Add("Content-Type", "application/json")

	response, _ := app.Test(request)
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)

	assert.Equal(t, `{"status":"OK"}`, string(body))	
}  



/* func TestWhenBadRequestIsGiven(t *testing.T){
	config.ConnectDB()
	app := todo.CreateApp()
	database := todo.NewDatabase()
	service := todo.NewService(database)
	handler := todo.NewHandler(service)

	handler.RegisterRoutes(app)

	requestBody := strings.NewReader("bad request")
	request := httptest.NewRequest(http.MethodPost, "/newtodo", requestBody)
	request.Header.Add("Content-Type", "application/json")

	response, _ := app.Test(request)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, `{"status":"Bad Request"}`, string(body) )
} */

func TestWhenBadRequestIsGiven(t *testing.T){
	app := todo.CreateApp()
	ctrl := gomock.NewController(t)
	mockService:=mocks.NewMockHandlerService(ctrl)
	
	handler := todo.NewHandler(mockService)

	handler.RegisterRoutes(app)
	requestBody := strings.NewReader("bad request")
	request := httptest.NewRequest(http.MethodPost, "/newtodo", requestBody)
	request.Header.Add("Content-Type", "application/json")

	response, _ := app.Test(request)
	body, _ := ioutil.ReadAll(response.Body)
	
	assert.Equal(t, `{"status":"Bad Request"}` ,string(body))

}
/* func TestAddToDo(t *testing.T){
	config.ConnectDB()
	app := todo.CreateApp()
	database := todo.NewDatabase()
	service := todo.NewService(database)
	handler := todo.NewHandler(service)

	handler.RegisterRoutes(app)

	
	requestStruct := todo.AddToDoRequest{
		Todo: "do your task",
	}
	expectedResponse := todo.AddToDoResponse{
		Todo: requestStruct.Todo,
		Message: "New ToDo Added to Databese",
	}
	body, _ := json.Marshal(&requestStruct)
	request := httptest.NewRequest(http.MethodPost, "/newtodo", bytes.NewBuffer(body))
	request.Header.Add("Content-Type", "application/json")

	response, _ := app.Test(request)
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, http.StatusCreated, response.StatusCode)

	var bodyStruct todo.AddToDoResponse 
	err = json.Unmarshal([]byte(string(body)), &bodyStruct)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}
	

	assert.Equal(t, expectedResponse, bodyStruct)
}
 */
 func TestAddToDo(t *testing.T){
	app := todo.CreateApp()
	ctrl := gomock.NewController(t)
	mockService:=mocks.NewMockHandlerService(ctrl)
	
	handler := todo.NewHandler(mockService)
	handler.RegisterRoutes(app)

	input := todo.AddToDoRequest{
		Todo: "last change",
	}
	output := todo.AddToDoResponse{
		Message: "New ToDo Added to Databese",
		Todo: input.Todo,
	}
	mockService.EXPECT().AddNewToDo(gomock.Any(), &input).Return(&output,nil)

	body, _ := json.Marshal(&input)

	request := httptest.NewRequest(http.MethodPost, "/newtodo", bytes.NewBuffer(body))
	request.Header.Add("Content-Type", "application/json")

	response, _ := app.Test(request)
	bodyResponse, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, http.StatusCreated, response.StatusCode)

	var actualResponse todo.AddToDoResponse 
	err := json.Unmarshal(bodyResponse, &actualResponse)

	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}

	assert.Equal(t, output, actualResponse )
 }

 func TestGetAllForHandler(t *testing.T){
	app := todo.CreateApp()
	ctrl := gomock.NewController(t)
	mockService:=mocks.NewMockHandlerService(ctrl)
	
	handler := todo.NewHandler(mockService)
	handler.RegisterRoutes(app)
	list := []todo.DatabaseStruct{}
	list = append(list, 
		todo.DatabaseStruct{
			ID:   [12]byte{},
			Todo: "omg!",
		},
		todo.DatabaseStruct{
			ID:   [12]byte{},
			Todo: "rtfm!",
		})
	output := todo.GetAllResponse{
		Todolist: list,
		Message: "Success",
	}
	mockService.EXPECT().GetAll().Return(&output)

	request := httptest.NewRequest(http.MethodGet, "/alltodo",http.NoBody)

	response,_ := app.Test(request)

	body, _ := ioutil.ReadAll(response.Body)

	var actualResponse todo.GetAllResponse
	err :=json.Unmarshal(body, &actualResponse)
	if err != nil{
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())

	}
	assert.Equal(t, output, actualResponse)

 }