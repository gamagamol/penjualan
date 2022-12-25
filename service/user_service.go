package service

import (
	"fmt"
	"log"
	"net/http"
	"sales/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
)



func (s *service)GetUser()(*[]entity.User,error){
	return s.r.GetUser()	
}

func (s *service)UserInsert(data entity.UserRequest)(entity.UserResponse,error){
	
	// hash password
	pass,_:=bcrypt.GenerateFromPassword([]byte(data.Password),bcrypt.DefaultCost)

	user:=entity.User{
		Name: data.Name,
		Ussername: data.Ussername,
		Password: string(pass),
		CreatedAt: time.Now(),
	}

	res,err:=s.r.UserInsert(user)

	if err!=nil{
		fmt.Println(err)
	}

	return res,err

}

func (s *service)UserUpdate(id int,data entity.UserRequest)(entity.UserResponse,error){

	req:=entity.User{
		IdUser: id,
		Name: data.Name,
		Ussername: data.Ussername,
		Password: data.Password,
		UpdatedAt: time.Now(),
	}

	

	res,err:=s.r.UserUpdate(id,req)

	if err!=nil{
		log.Println(err)
	}

	return res,err



}

func (s *service) UserGetById(id int,req entity.LoginRequest)(entity.UserResponse,error){

	if req.Ussername ==""{
	return s.r.UserGetById(id,"")

	}else{

		data,err:=s.r.UserGetById(id,req.Ussername)
		// check ussername
		fmt.Println(err)
		if err !=nil{
			return entity.UserResponse{
				Status: http.StatusBadRequest,
				Message: "Wrong ussername",
			},err
		}
		// check password
		err=bcrypt.CompareHashAndPassword([]byte(data.Login.Password),[]byte(req.Password))		
		if err!=nil{
			return entity.UserResponse{
				Status: http.StatusBadRequest,
				Message: "Wrong Password or Email",
			},err
		}

		return data,nil

	}
}

func (s *service)UserDelete(id int)(entity.UserResponse,error){

	req,err:=s.r.UserGetById(id,"")
	if err !=nil {
		fmt.Println(err)
	}

	data:=entity.User{
		IdUser: id,
		Name: req.Detail.Name,
		Ussername: req.Detail.Ussername,
		Password: req.Detail.Password,
		DeletedAt: time.Now(),
	}
	return s.r.UserDelete(id,data)
}


func (s *service)Test()string{
return "Hello World"
}