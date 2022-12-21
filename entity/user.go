package entity

import "time"

type User struct{

	IdUser 		int
	Name 		string
	Ussername 	string
	Password 	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt	time.Time
}

type UserRequest struct{
	Name string			`json:"name" binding:"required"`
	Ussername string	`json:"ussername" binding:"required"`
	Password string		`json:"password" binding:"required"`
}



