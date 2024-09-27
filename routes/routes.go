package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Search
	server.POST("/search", search)

	// Add search query
	server.POST("/search/add", addSearchQuery)

	// Delete search query
	server.DELETE("/search/:searchQuery", deleteSearchQuery)
}