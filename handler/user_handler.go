package handler

import (
	"fmt"
	"net/http"
	"sales/entity"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


func(h handler)GetUser(c *gin.Context){

	data,err:=h.s.GetUser()
	if err!=nil{
		fmt.Println(err)
	}
	c.JSON(http.StatusOK,data)
}


func (h handler)InsertUser(c *gin.Context){
	var user entity.UserRequest
	err:=c.ShouldBindJSON(&user)

	if err!=nil{
		
		errorMessages:=[]string{}
		for _,e:=range err.(validator.ValidationErrors){
			errorMessage:=fmt.Sprintf("Error on field %s, must be %s",e.Field(),e.ActualTag())
			errorMessages= append(errorMessages,errorMessage)
		}

		c.JSON(http.StatusBadRequest,errorMessages)
	}

	res,err:=h.s.UserInsert(user)

	if err !=nil{
		fmt.Println(err)
	}

	c.JSON(http.StatusAccepted,res)
}



func (h handler)UserGetById(c *gin.Context){
	id,_:=strconv.Atoi(c.Param("id"))

		res,err:=h.s.UserGetById(id,entity.LoginRequest{})

	if err!=nil{
		fmt.Println(err)
	}

	c.JSON(http.StatusOK,res)


}

func (h handler)Login(c *gin.Context){
	ussername:=c.Query("ussername")
	password:=c.Query("password")

	req:=entity.LoginRequest{
		Ussername: ussername,
		Password: password,
	}

	res,err:=h.s.UserGetById(0,req)

	if err!=nil{
		fmt.Println(err)
	}

	c.SetCookie("Aut",*res.Token,time.Now().Minute(),"","",false,true)

	c.JSON(http.StatusOK,res)
}

func (h handler)UserUpdate(c *gin.Context){
	id,_:=strconv.Atoi(c.Query("id"))


	

	res,err:=h.s.UserUpdate(id,entity.UserRequest{
		Name: c.Query("name"),
		Ussername: c.Query("ussername"),
		Password: c.Query("password"),
	})

	if err!=nil{
		fmt.Println(err)
	}

	c.JSON(http.StatusOK,res)

}


func (h handler)UserDelete(c *gin.Context){
	id,_:=strconv.Atoi(c.Query("id"))


	res,err:=h.s.UserDelete(id)

	if err !=nil{
		fmt.Println(err)
	}

	c.JSON(http.StatusOK,res)

}



