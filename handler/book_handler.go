package handler

import (
	"fmt"
	"log"
	"net/http"
	"sales/entity"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)





func (b *handler) GetAll(c *gin.Context){

	books,err:=b.s.GetAll()

	if err !=nil{

		log.Println("err")
	}


	c.JSON(http.StatusOK,books)
}

func (b *handler) Insert(c *gin.Context){

	var BookRequest entity.BookRequest
	err:=c.ShouldBindJSON(&BookRequest)
	
	if err!=nil{
		
		errorMessages:=[]string{}
		for _,e:=range err.(validator.ValidationErrors){
		 fmt.Println(e.Field())
			errorMessage:=fmt.Sprintf("Error on field %s, must be %s",e.Field(),e.ActualTag())
			errorMessages= append(errorMessages,errorMessage)
		}

		c.JSON(http.StatusBadRequest,errorMessages)
	}

	data,err:=b.s.Insert(BookRequest)

	if err!=nil{
		fmt.Println(err)
	}

	c.JSON(http.StatusOK,data)
	
}


func (b *handler) GetById(c *gin.Context){
	id,_:=strconv.Atoi(c.Param("id"))

	book,err := b.s.GetById(id)

	if err!=nil{
		fmt.Println(err)
	}

	c.JSON(http.StatusOK,book)
}


func (b *handler) Update(c *gin.Context){
	id,_:=strconv.Atoi(c.Query("id"))
	title:=c.Query("title")
	price,_:=strconv.Atoi(c.Query("price"))
	stock,_:=strconv.Atoi(c.Query("stock"))

	book:=entity.BookRequest{
		Title: title,
		Price: price,
		Stock: stock,
	}


	bookResponse,err:=b.s.Update(id,book)

	if err !=nil{
		fmt.Println(err)
	}
	c.JSON(http.StatusOK,bookResponse)

}


func (b handler) Delete(c *gin.Context){
	id,_:=strconv.Atoi(c.Query("id"))

	book,err:=b.s.Delete(id)

	if err!=nil{
		fmt.Println(err)
	}

	c.JSON(http.StatusOK,book)
}