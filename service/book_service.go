package service

import (
	// "fmt"
	"fmt"
	"net/http"
	"sales/entity"
	"time"
)




func (s *service) GetAll()([]entity.Book,error){
	
	return s.r.GetAll()
}


func (s *service)Insert(data entity.BookRequest)(entity.BookResponse,error){
	var book entity.Book

	book.Title=data.Title
	book.Price=data.Price
	book.Stock=data.Stock
	book.CreatedAt=time.Now()

	return s.r.Insert(book)
	
}

func (s *service) GetById(id int)(entity.BookResponse,error){
	return s.r.GetbyId(id)
}
func (s *service)Update(id int,data entity.BookRequest)(entity.BookResponse,error){


	BookData:=entity.Book{
		IdBook: id,
		Title: data.Title,
		Price: data.Price,
		Stock: data.Stock,
		UpdatedAt: time.Now(),
	}

	return s.r.Update(BookData)


}

func (s *service)Delete(id int)(*entity.BookResponse,error){


	bookid,err:=s.GetById(id)
	if err !=nil{
		fmt.Println(err)
	}

	data:=entity.Book{
		IdBook: id,
		Title: bookid.Detail.Title,
		Price: bookid.Detail.Price,
		Stock: bookid.Detail.Stock,
		DeletedAt: time.Now(),
	}
	
	err=s.r.Delete(id,data)
	if err!=nil{
		fmt.Println(err)
	}
	book:=entity.BookResponse{
		Status: http.StatusOK,
		Message: "Book has Been delete",
	}

	return &book,err


}

