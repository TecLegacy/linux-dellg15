package main

import (
	"employeeTestify/internal/config"
	"employeeTestify/internal/router"
)

func main() {
	config.ConnectDatabase()
	r := router.SetupRouter()
	r.Run(":8080")
}
