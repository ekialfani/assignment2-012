package main

import (
	"assignment2-012/routers"
)

func main() {
	routers.StartServer().Run(":8080")
}