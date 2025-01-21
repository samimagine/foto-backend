package routes

import (
	"foto-backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:id", controllers.GetUser)    // Get user by ID
		userGroup.POST("/", controllers.CreateUser)   // Create user
		userGroup.PUT("/:id", controllers.UpdateUser) // Update user
		userGroup.DELETE("/:id", controllers.DeleteUser) // Delete user
	}
}