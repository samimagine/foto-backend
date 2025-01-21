package controllers

import (
	"foto-backend/config"
	"foto-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindNearbyUsers(c *gin.Context) {
	latitude, _ := strconv.ParseFloat(c.Query("latitude"), 64)
	longitude, _ := strconv.ParseFloat(c.Query("longitude"), 64)
	radius, _ := strconv.ParseFloat(c.Query("radius"), 64)

	var users []models.User
	config.DB.Raw(`
		SELECT *, (
			6371 * acos(
				cos(radians(?)) * cos(radians(latitude)) *
				cos(radians(longitude) - radians(?)) +
				sin(radians(?)) * sin(radians(latitude))
			)
		) AS distance
		FROM users
		HAVING distance <= ?
		ORDER BY distance ASC
	`, latitude, longitude, latitude, radius).Scan(&users)

	c.JSON(http.StatusOK, users)
}