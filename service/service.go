package service

import (
	"sales/entity"
	"sales/repository"
)

type Service interface{
	GetAll()([]entity.Book,error)
	GetById(id int)(entity.BookResponse,error)
	Insert(data entity.BookRequest)(entity.BookResponse,error)
	Update(id int, data entity.BookRequest)(entity.BookResponse,error)
	Delete(id int)(*entity.BookResponse,error)
	GetUser() (*[]entity.User,error)
	UserGetById(id int)(entity.UserResponse,error)
	UserInsert(data entity.UserRequest)(entity.UserResponse,error)
	UserUpdate(id int, data entity.UserRequest)(entity.UserResponse,error)
	UserDelete(id int,)(entity.UserResponse,error)

}

type service struct{
	r 	repository.Repository
	
}


func Newservice(r repository.Repository)*service{
	return &service{r}
}