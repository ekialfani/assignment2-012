package controllers

import (
	"assignment2-012/database"
	"assignment2-012/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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