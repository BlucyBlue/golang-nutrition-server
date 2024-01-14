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
	router.POST("/products", AddProductEndpoint)
	router.GET("/products/:productID", GetProductByIDEndpoint)
	router.PUT("/products", UpdateProductEndpoint)
	router.DELETE("/products/:productID", DeleteProductEndpoint)
	return router
}

func main() {
	router := SetupRouter()
	router.Run()
}
