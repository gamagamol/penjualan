package entity

type BookResponse struct{
	Status int `json:"status"`
	Message string `json:"message"`
	Detail BookRequest
}


type UserResponse struct{
	Status int `json:"status"`
	Message string `json:"message"`
	Detail *UserRequest
}


