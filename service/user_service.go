package service

import (
	"fmt"
	"log"
	"sales/entity"
	"time"
)



func (s *service)GetUser()(*[]entity.User,error){
	return s.r.GetUser()	
}

func (s *service)UserInsert(data entity.UserRequest)(entity.UserResponse,error){
	
	user:=entity.User{
		Name: data.Name,
		Ussername: data.Ussername,
		Password: data.Password,
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

func (s *service) UserGetById(id int)(entity.UserResponse,error){

	return s.r.UserGetById(id)
}

func (s *service)UserDelete(id int)(entity.UserResponse,error){

	req,err:=s.UserGetById(id)
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