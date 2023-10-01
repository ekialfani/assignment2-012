package database

import (
	"github.com/ekialfani/assignment2-012/config"
	"github.com/ekialfani/assignment2-012/models"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)

func StartDB() {
	config := config.GetDBConfig()

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error:", err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}