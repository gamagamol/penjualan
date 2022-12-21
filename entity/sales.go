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