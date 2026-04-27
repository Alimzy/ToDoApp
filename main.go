package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.Default() gives us a router with Logger and Recovery middleware
	r := gin.Default()

	// A simple health-check route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Start the server on port 8080
	r.Run(":8080")
}
