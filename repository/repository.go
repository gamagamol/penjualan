package repository

import (
	"sales/entity"

	"gorm.io/gorm"
)

type Repository interface{
	// book repo
	GetAll()([]entity.Book,error)
	GetbyId(id int)(entity.BookResponse,error)
	Insert(data entity.Book)(entity.BookResponse,error)
	Delete(id int,data entity.Book) error
	Update(dat entity.Book)(entity.BookResponse,error)
	// user repo
	GetUser()(*[]entity.User,error)
	UserGetById(id int,ussername string)(entity.UserResponse,error)
	UserInsert(data entity.User)(entity.UserResponse,error)
	UserUpdate(id int, data entity.User)(entity.UserResponse,error)
	UserDelete(id int,data entity.User)(entity.UserResponse,error)
	// sales repo
	GetSales()(entity.SalesResponse,error)
	GetSalesById(id int)(entity.SalesResponse,error)
	SalesInsert(data entity.Sales)(entity.SalesResponse,error)
	SalesUpdate(id int,data entity.Sales)(entity.SalesResponse,error)

}

type repostory struct{
	db	*gorm.DB
}


func Newrepostory(db *gorm.DB)*repostory{
	return &repostory{db}
}