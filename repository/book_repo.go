package repository

import (
	"fmt"
	"net/http"
	"sales/entity"
)







func (r *repostory) GetAll()([]entity.Book,error){
	var book []entity.Book
	err:=r.db.Where("deleted_at IS NULL").Find(&book).Error
	return book,err
}

func (r *repostory) Insert(data entity.Book)(entity.BookResponse,error){
	
	book:=entity.BookResponse{
		Status: http.StatusCreated,
		Message: fmt.Sprintf("Success Insert New Book (%s)",data.Title),
		Detail: entity.BookRequest{
			Title: data.Title,
			Price: data.Price,
			Stock: data.Stock,	
		},

	}
	err:=r.db.Omit("updated_at","deleted_at").Create(&data).Error
	return book,err
}

func (r *repostory)GetbyId(id int)(entity.BookResponse,error){
	var book entity.Book
	
	err:= r.db.Find(&book,id).Error

	bookResponse:=entity.BookResponse{
		Status: http.StatusOK,
		Message: "Book Finded",
		Detail: entity.BookRequest{
			Title: book.Title,
			Price: book.Price,
			Stock: book.Stock,
		},
	}
	return bookResponse,err
}

func (r *repostory)Update(data entity.Book)(entity.BookResponse,error){

	err:=r.db.Omit("created_at","deleted_at").Where("id_book",data.IdBook).Save(&data).Error
	book:=entity.BookResponse{
		Status: http.StatusOK,
		Message: "Book Has been Updated",
		Detail: entity.BookRequest{
			Title: data.Title,
			Price: data.Price,
			Stock: data.Stock,
		},
	}
	return book,err
}

func (r *repostory)Delete(id int,data entity.Book) error{

	err:=r.db.Where("id_book",id).Omit("created_at","updated_at").Save(&data).Error

	return err
}






