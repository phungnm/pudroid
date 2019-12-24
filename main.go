package main

import (
	"github.com/gin-gonic/gin"
	"api/controllers"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/public", "./public")

	client := router.Group("/api")
	{
		client.GET("/story/detail/:id", controllers.Read)
		client.GET("/story/getAll", controllers.GetAll)
		client.POST("/story/create", controllers.Create)
		client.PATCH("/story/update/:id", controllers.Update)
		client.DELETE("/story/:id", controllers.Delete)
	}
	
	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000") // Ứng dụng chạy tại cổng 8080
}