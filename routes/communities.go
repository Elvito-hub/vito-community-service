package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vitocom.com/community/models"
)

func getAllCommunities(context *gin.Context) {
	communities, err := models.GetAllCommunities()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch"})
	}

	context.JSON(http.StatusOK, communities)
}
