package controllers

import (
	"assignment2-012/database"
	"assignment2-012/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOrderById(context *gin.Context) {
	db := database.GetDB()
	id := context.Param("id")

	var order models.Order

	err := db.Preload("Items").First(&order, id).Error

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_message": fmt.Sprintf("user with id %v not found", id),
		})
		return
	}

	for _, item := range order.Items {
		db.Delete(&item)
	}

	db.Delete(&order)

	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("order with id %v has been successfully deleted", id),
	})
}

func UpdateOrderById(context *gin.Context) {
	db := database.GetDB()
	id := context.Param("id")

	var order models.Order

	err := db.Preload("Items").First(&order, id).Error

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_message": fmt.Sprintf("user with id %v not found", id),
		})
		return
	}

	if err := context.ShouldBindJSON(&order); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	for _, item := range order.Items {
		db.Save(&item)
	}

	db.Save(&order)

	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": order,
	})
}

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