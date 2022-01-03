package todo


type AddToDoResponse struct{
	Message string `json:"message"` 
	Todo string `json:"todo"`
}

type GetAllResponse struct{
	Todolist []DatabaseStruct `json:"todolist"`
	Message string `json:"message"`
}