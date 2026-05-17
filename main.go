package main

import (
	"log"

	"github.com/joho/godotenv"
	"gotask/config"
	"gotask/db"
	"gotask/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment")
	}

	cfg := config.Load()
	db.Connect(cfg)
	r := routes.SetupRoutes()
	r.Run(":8080")
}
