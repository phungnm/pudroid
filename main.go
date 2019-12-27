package main

import (
	"github.com/gin-gonic/gin"
	"pudroid/controllers"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/public", "./public")

	client := router.Group("/api/shortenUrl")
	{
		  client.GET("/get", controllers.GetShortenUrl)
		  client.POST("/add", controllers.AddShortenUrl)
		// client.GET("/story/getAll", controllers.GetAll)
		// client.POST("/story/create", controllers.Create)
		// client.PATCH("/story/update/:id", controllers.Update)
		// client.DELETE("/story/:id", controllers.Delete)
	}
	
	return router
}

func main() {
  	router := setupRouter()
	router.Run(":3000") // Ứng dụng chạy tại cổng 8080
}