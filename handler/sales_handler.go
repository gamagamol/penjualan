package handler

import (
	"fmt"
	"net/http"
	"sales/entity"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// test
func (h handler)GetSales(c *gin.Context){
	res,err:=h.s.GetSales()
	if err!=nil{
		fmt.Println(err)
	}
	c.JSON(http.StatusOK,res)
}
func (h handler)GetSalesById(c *gin.Context){
	id,_:=strconv.Atoi(c.Param("id"))
	res,err:=h.s.GetSalesById(id)
	if err!=nil{
		fmt.Println(err)
	}
	c.JSON(http.StatusOK,res)
}


func (h handler)SalesInsert(c *gin.Context){
	var sales entity.SalesRequest

	err:=c.ShouldBindJSON(&sales)

	if err!=nil{
		
		errorMessages:=[]string{}
		for _,e:=range err.(validator.ValidationErrors){
			errorMessage:=fmt.Sprintf("Error on field %s, must be %s",e.Field(),e.ActualTag())
			errorMessages= append(errorMessages,errorMessage)
		}

		c.JSON(http.StatusBadRequest,errorMessages)
	}

	res,er:=h.s.SalesInsert(sales)

	if er!=nil{
		fmt.Println(er)
	}

	c.JSON(http.StatusCreated,res)

}


func (h handler)SalesUpdate(c *gin.Context){

	id,_:=strconv.Atoi(c.Query("id"))
	id_book,_:=strconv.Atoi(c.Query("id_book"))
	id_user,_:=strconv.Atoi(c.Query("id_user"))
	price,_:=strconv.Atoi(c.Query("price"))
	unit,_:=strconv.Atoi(c.Query("unit"))
	total,_:=strconv.Atoi(c.Query("total"))

	req:=entity.SalesRequest{
		IdBook: id_book,
		IdUser: id_user,
		Price: price,
		Unit: unit,
		Total: total,
	}

	data,err:=h.s.SalesUpdate(id,req)

	if err!=nil{
		fmt.Println(err)
	}

	c.JSON(http.StatusOK,data)



}