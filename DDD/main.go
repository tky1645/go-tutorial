package main

import (
	"DDD/command/user"

	"github.com/gin-gonic/gin"
)


func main (){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "ping pong")
	})
	r.GET("/users", user.HandlerGET)
	r.POST("/users", user.HandlerPOST)

	r.Run(":8080")
}