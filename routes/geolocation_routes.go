package routes

import (
	"foto-backend/controllers"
	"github.com/gin-gonic/gin"
)

func GeolocationRoutes(router *gin.Engine) {
	router.GET("/users/nearby", controllers.FindNearbyUsers)
}