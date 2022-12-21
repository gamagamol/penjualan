package entity

import (
	"time"
)

type Book struct{
	IdBook int				
	Title string		
	Price int
	Stock int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type BookRequest struct{
	Title string `json:"title" binding:"required"`
	Price int `json:"price" binding:"required"`
	Stock int `json:"stock" binding:"required"`
}


