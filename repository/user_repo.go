package repository

import (
	"errors"
	"fmt"
	"net/http"
	"sales/entity"

	"gorm.io/gorm"
)

func(r *repostory) GetUser()(*[]entity.User,error){
	var User []entity.User
	
	err:=r.db.Where("deleted_at IS NULL").Omit("created_at","updated_at","deleted_at").Find(&User).Error
	return &User,err
}

func (r *repostory)UserGetById(id int,ussername string)(entity.UserResponse,error){
	var user entity.User
	var err error

	if ussername==""{
		err=r.db.Find(&user,id).Error
			if err!=nil{
				fmt.Println("error di repo")
				fmt.Println(err)
			}	

	return entity.UserResponse{
		Status: http.StatusOK,
		Message: "ok",
		Detail: &entity.UserRequest{
			Name: user.Name,
			Ussername: user.Ussername,
			Password: user.Password,
		},
	},err
	}else{

		err=r.db.Where("ussername",ussername).Find(&user).Error
			if errors.Is(err,gorm.ErrRecordNotFound){

				return entity.UserResponse{},err
			}
			
	return entity.UserResponse{
		Status: http.StatusOK,
		Message: "ok",
		Login: &entity.LoginRequest{
			Ussername: user.Ussername,
			Password: user.Password,
		},
	},err	
	}

	


}

func (r *repostory)UserInsert(data entity.User)(entity.UserResponse,error){
	
	err:= r.db.Omit("updated_at","deleted_at").Create(&data).Error
	if err!=nil{
		fmt.Println(err)
	}

	res:=entity.UserResponse{
		Status: http.StatusCreated,
		Message: "Success Insert New User",
		Detail: &entity.UserRequest{
			Name: data.Name,
			Ussername: data.Ussername,
			Password: data.Password,
		},
	}

	return res,err
}

func (r *repostory)UserUpdate(id int,data entity.User)(entity.UserResponse,error){
	err:=r.db.Where("id_user",id).Omit("created_at","deleted_at").Save(&data).Error

	res:=entity.UserResponse{
		Status: http.StatusOK,
		Message: "Success Update User",
		Detail: &entity.UserRequest{
			Name: data.Name,
			Ussername: data.Ussername,
			Password: data.Password,
		},
	}
	return res,err
}


func (r *repostory)UserDelete(id int,req entity.User)(entity.UserResponse,error){
	
	err:=r.db.Where("id_user",id).Omit("created_at","updated_at").Save(&req).Error

	detail,er:=r.UserGetById(id,"")

	if er !=nil{
		fmt.Println(er)
	}

	return entity.UserResponse{
		Status: http.StatusOK,
		Message: "Success Delete User",
		Detail: detail.Detail,
	},err

}