package repository

import (
	"fmt"
	"net/http"
	"sales/entity"
)


func (r *repostory)GetSales()(entity.SalesResponse,error){
	
		var sales []entity.SalesJoinResponse
		err:=r.db.Raw("select date_sales,title,name,sales.price,unit,total from sales join books on sales.id_book=books.id_book join users on users.id_user=sales.id_user").Scan(&sales).Error
		salesResponse:=entity.SalesResponse{
			Status: http.StatusOK,
			Message: "ok",
			Data: &sales,
			
		}
		return salesResponse,err
}

func (r *repostory)SalesInsert(data entity.Sales)(entity.SalesResponse,error){
	
	err:=r.db.Create(&data).Error
	if err!=nil{
		fmt.Println(err)
	}
	return entity.SalesResponse{
		Status: http.StatusOK,
		Message: "ok",
		Detail: &entity.SalesRequest{
			IdBook: data.IdBook,
			IdUser: data.IdUser,
			Price: data.Price,
			Unit: data.Unit,
			Total: data.Total,
		} ,
	},err
}

func (r *repostory)SalesUpdate(id int ,data entity.Sales)(entity.SalesResponse,error){

	err:=r.db.Where("id_sales",id).Save(&data).Error

	if err!=nil{
		fmt.Println(err)
	}

	return entity.SalesResponse{
		Status: http.StatusOK,
		Message: "ok",
		Detail: &entity.SalesRequest{
			IdBook: data.IdBook,
			IdUser: data.IdUser,
			Price: data.Price,
			Unit: data.Unit,
			Total: data.Total,
		},
	},err


}
func (r *repostory)GetSalesById(id int)(entity.SalesResponse,error){
	
		var sales entity.SalesRequest
		err:=r.db.Raw("select date_sales,title,name,sales.price,unit,total from sales join books on sales.id_book=books.id_book join users on users.id_user=sales.id_user").Where("id_sales",id).Scan(&sales).Error
		salesResponse:=entity.SalesResponse{
			Status: http.StatusOK,
			Message: "ok",
			Detail: &sales,
			
		}
		return salesResponse,err
}

