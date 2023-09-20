package controllers

import (
	"assignment2-012/database"
	"assignment2-012/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(context *gin.Context) {
	db := database.GetDB()
	
	var orders []models.Order

	err := db.Preload("Items").Find(&orders).Error

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_message": "Failed to retrieve orders data",
			"error_detail":  err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"Data": orders,
	})
}

func CreateOrder(context *gin.Context) {
	db := database.GetDB()

	var order models.Order
	if err := context.ShouldBindJSON(&order); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Create(&order).Error

	if err != nil {
		context.JSON(500, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"order": order,
	})
}