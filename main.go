package main

import (
	"github.com/ekialfani/assignment2-012/routers"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	routers.StartServer().Run(":8080")
}