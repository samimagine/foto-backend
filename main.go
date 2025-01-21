package main

import (
    "foto-backend/config"
    "foto-backend/models"
	"foto-backend/routes"
)

func main() {
    config.ConnectDatabase()

    config.DB.AutoMigrate(&models.User{}, &models.Rating{})
	r := routes.SetupRoutes()
    r.Run(":8080") // Start the server
}