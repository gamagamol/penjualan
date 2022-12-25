package service

import (
	"sales/entity"
	"sales/repository"
)

type Service interface{
	// book service
	GetAll()(entity.BookResponse,error)
	GetById(id int)(entity.BookResponse,error)
	Insert(data entity.BookRequest)(entity.BookResponse,error)
	Update(id int, data entity.BookRequest)(entity.BookResponse,error)
	Delete(id int)(*entity.BookResponse,error)
	// user service
	GetUser() (*[]entity.User,error)
	UserGetById(id int,req entity.LoginRequest)(entity.UserResponse,error)
	UserInsert(data entity.UserRequest)(entity.UserResponse,error)
	UserUpdate(id int, data entity.UserRequest)(entity.UserResponse,error)
	UserDelete(id int,)(entity.UserResponse,error)
	// Login(req entity.UserRequest)(entity.BookResponse,error)

	// sales service
	GetSales()(entity.SalesResponse,error)
	GetSalesById(id int)(entity.SalesResponse,error)
	SalesInsert(req entity.SalesRequest)(entity.SalesResponse,error)
	SalesUpdate(id int,data entity.SalesRequest)(entity.SalesResponse,error)


	Test()string

}

type service struct{
	r 	repository.Repository
	
}


func Newservice(r repository.Repository)*service{
	return &service{r}
}