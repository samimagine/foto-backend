package routes

import (
    "foto-backend/controllers"

    "github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
    r := gin.Default()

    // Authentication
    r.POST("/signup", controllers.Signup)
    r.POST("/login", controllers.Login)

    return r
}