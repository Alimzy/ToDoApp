package main

import (
	"gotask/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.Default() gives us a router with Logger and Recovery middleware
	r := gin.Default()

	// A simple health-check route
	routes.SetupRoutes(r)

	// Start the server on port 8080
	r.Run(":8080")
}
