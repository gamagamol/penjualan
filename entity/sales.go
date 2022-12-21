package entity

import "time"

type Sales struct{
	IdSales int 
	DateSales time.Time
	IdBook int
	IdUser int
	Price int
	Unit int
	Total int
}

type SalesRequest struct{
	IdBook 	int	`json:"id_book,omitempty" binding:"required"` 
	IdUser 	int	`json:"id_user,omitempty" binding:"required"`
	Price 	int	`json:"price,omitempty" binding:"required"`
	Unit 	int	`json:"unit,omitempty" binding:"required"`
	Total 	int	`json:"total,omitempty" binding:"required"`
}