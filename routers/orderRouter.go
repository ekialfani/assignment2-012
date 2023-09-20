package routers

import (
	"assignment2-012/database"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	database.StartDB()

	router := gin.Default()

	router.GET("/test", nil)

	return router
}