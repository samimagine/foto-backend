package routes

import (
	"foto-backend/controllers"
	"github.com/gin-gonic/gin"
)

func ChatRoutes(router *gin.Engine) {
	router.GET("/chat", controllers.ChatHandler)
}