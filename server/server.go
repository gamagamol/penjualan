package server

import (
	"log"
	"sales/handler"
	"sales/middleware"
	"sales/repository"
	"sales/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartServe() {

	dsn := "root:@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	Repository := repository.Newrepostory(db)
	Service := service.Newservice(Repository)
	handler := handler.NewHandler(Service)

	router := gin.Default()
	// book Handler
	router.GET("/book", middleware.Auth, handler.GetAll)
	router.GET("/book/:id", middleware.Auth, handler.GetById)
	router.POST("/book/insert", middleware.Auth, handler.Insert)
	router.PUT("/book/update", middleware.Auth, handler.Update)
	router.DELETE("/book/delete", middleware.Auth, handler.Delete)

	// user Handler
	router.GET("/user", middleware.Auth, handler.GetUser)
	router.GET("/user/:id", middleware.Auth, handler.UserGetById)
	router.GET("/user/login", handler.Login)
	router.POST("/user/insert", middleware.Auth, handler.InsertUser)
	router.PUT("/user/update", middleware.Auth, handler.UserUpdate)
	router.DELETE("/user/delete", middleware.Auth, handler.UserDelete)

	// sales handler
	router.GET("/sales", middleware.Auth, handler.GetSales)
	router.GET("/sales/:id", middleware.Auth, handler.GetSalesById)
	router.POST("/sales/insert", middleware.Auth, handler.SalesInsert)
	router.PUT("/sales/update", middleware.Auth, handler.SalesUpdate)

	router.Run(":80")
}
