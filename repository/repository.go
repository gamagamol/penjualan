package repository

import (
	"sales/entity"

	"gorm.io/gorm"
)

type Repository interface{
	GetAll()([]entity.Book,error)
	GetbyId(id int)(entity.BookResponse,error)
	Insert(data entity.Book)(entity.BookResponse,error)
	Delete(id int,data entity.Book) error
	Update(dat entity.Book)(entity.BookResponse,error)
	GetUser()(*[]entity.User,error)
	UserGetById(id int)(entity.UserResponse,error)
	UserInsert(data entity.User)(entity.UserResponse,error)
	UserUpdate(id int, data entity.User)(entity.UserResponse,error)
	UserDelete(id int,data entity.User)(entity.UserResponse,error)

}

type repostory struct{
	db	*gorm.DB
}


func Newrepostory(db *gorm.DB)*repostory{
	return &repostory{db}
}