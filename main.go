package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is up!",
		})
	})
	router.POST("/register", RegisterUser)
	return router
}

func main() {
	router := SetupRouter()
	router.Run()
}
