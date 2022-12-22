package service

import (
	"fmt"
	"sales/entity"
	"time"
)

// test
func (s *service)GetSales()(entity.SalesResponse,error){
	return s.r.GetSales()
}

func (s *service)SalesInsert(req entity.SalesRequest)(entity.SalesResponse,error){
	
	data:=entity.Sales{
		IdBook: req.IdBook,
		DateSales: time.Now(),
		IdUser: req.IdUser,
		Price: req.Price,
		Unit: req.Unit,
		Total: req.Total,
		
	}


	res,err:=s.r.SalesInsert(data)
	if err!=nil{
		fmt.Println(err)
	}

	return res,err
}

func (s *service)SalesUpdate(id int,req entity.SalesRequest)(entity.SalesResponse,error){
	data:=entity.Sales{
		IdSales: id,
		DateSales: time.Now(),
		IdBook: req.IdBook,
		IdUser: req.IdUser,
		Price: req.Price,
		Unit: req.Unit,
		Total: req.Total,
	}

	return s.r.SalesUpdate(id,data)
}

func (s *service)GetSalesById(id int)(entity.SalesResponse,error){
	return s.r.GetSalesById(id)
}