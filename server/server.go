package server

import (
	"log"
	"sales/handler"
	"sales/repository"
	"sales/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartServe(){

	 dsn := "root:@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err!=nil{
		log.Println(err)
	}

	Repository:=repository.Newrepostory(db)
	Service:=service.Newservice(Repository)
	handler:=handler.NewHandler(Service)

	router:=gin.Default()
	// book Handler
	router.GET("/book",handler.GetAll)
	router.GET("/book/:id",handler.GetById)
	router.POST("/book/insert",handler.Insert)
	router.PUT("/book/update",handler.Update)
	router.DELETE("/book/delete",handler.Delete)

	// user Handler
	router.GET("/user",handler.GetUser)
	router.GET("/user/:id",handler.UserGetById)
	router.GET("/user/login",handler.Login)
	router.POST("/user/insert",handler.InsertUser)
	router.PUT("/user/update",handler.UserUpdate)
	router.DELETE("/user/delete",handler.UserDelete)

	// sales handler
	router.GET("/sales",handler.GetSales)
	router.GET("/sales/:id",handler.GetSalesById)
	router.POST("/sales/insert",handler.SalesInsert)
	router.PUT("/sales/update",handler.SalesUpdate)


	router.GET("/test",handler.Test)
	router.Run(":80")
}