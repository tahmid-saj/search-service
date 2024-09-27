package routes

import (
	"net/http"
	"search-service/models"

	"github.com/gin-gonic/gin"
)

func search(context *gin.Context) {
	var searchQueryInput models.SearchQueryInput

	err := context.ShouldBindJSON(&searchQueryInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	searchResults, err := models.Search(searchQueryInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch search results"})
		return
	}

	context.JSON(http.StatusOK, searchResults)
}

func addSearchQuery(context *gin.Context) {
	var searchQueryInput models.SearchQueryInput

	err := context.ShouldBindJSON(&searchQueryInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	searchQueryAdded, err := models.AddSearchQuery(searchQueryInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not add search query"})
		return
	}

	context.JSON(http.StatusCreated, searchQueryAdded)
}

func deleteSearchQuery(context *gin.Context) {
	var searchQueryInput models.SearchQueryInput

	err := context.ShouldBindJSON(&searchQueryInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	searchQueryDeleted, err := models.DeleteSearchQuery(searchQueryInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete search query"})
		return
	}

	context.JSON(http.StatusOK, searchQueryDeleted)
}