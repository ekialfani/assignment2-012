package main

import (
	"assignment2-012/database"
	"fmt"
)

func main() {
	fmt.Println("start connected to database...")
	database.StartDB()
}