package routers

import (
	"assignment2-012/controllers"
	"assignment2-012/database"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	database.StartDB()

	router := gin.Default()

	router.POST("/order", controllers.CreateOrder)
	router.GET("/orders", controllers.GetAllOrders)
	router.PUT("/orders/:id", controllers.UpdateOrderById)

	return router
}