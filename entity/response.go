package entity

import "time"

type BookResponse struct{
	Status 	int `json:"status"`
	Message string `json:"message"`
	Data	*[]Book
	Detail  *BookRequest
	
}


type UserResponse struct{
	Status int `json:"status"`
	Message string `json:"message"`
	Detail *UserRequest	`json:"detail,omitempty"`
	Login *LoginRequest `json:"login,omitempty"`
}

type SalesJoinResponse struct{
	IdSales int
	DateSales time.Time
	IdBook int
	IdUser int
	Price int
	Unit int
	Total int
}

type SalesResponse struct{
	Status int `json:"status"`
	Message string `json:"message"`
	Data *[]SalesJoinResponse `json:"data,omitempty"`
	Detail *SalesRequest `json:"detail,omitempty"`
}




