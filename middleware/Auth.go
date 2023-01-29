package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func Auth(c *gin.Context){


	cookie,err:=c.Cookie("Aut")
	var token string

	if c.Request.Header["Authorization"] == nil{
		
		c.JSON(http.StatusBadRequest,gin.H{
			"Message":"Please Fill Token",
		})
		token="test"
	}else{

		token=strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	}

	
	

	if err !=nil{
		 c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"Message":"token is invalid",
		})
	}

	if token==""{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"Message":"Please fill token",
		})
	}

	// check token is not  same
	if cookie !=token{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"Message":"token is invalid not same",
		})
	}


	c.Next()

}