package entity

import "time"

type User struct{

	IdUser 		int
	Name 		string `json:"name,omitempty"`
	Ussername 	string
	Password 	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time `json:"updatedt_at,omitempty"`
	DeletedAt	time.Time	`json:"deleted_at,omitempty"`
}

type UserRequest struct{
	Name string			`json:"name,omitempty" binding:"required"`
	Ussername string	`json:"ussername" binding:"required"`
	Password string		`json:"password" binding:"required"`
}


type LoginRequest struct{
	Ussername string	`json:"ussername" binding:"required"`
	Password string		`json:"password" binding:"required"`
}



