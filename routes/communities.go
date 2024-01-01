package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vitocom.com/community/models"
)

func getAllCommunities(context *gin.Context) {
	communities, err := models.GetAllCommunities()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch"})
	}

	context.JSON(http.StatusOK, communities)
}

func createCommunity(context *gin.Context) {
	var com models.Community

	err := context.ShouldBindJSON(&com)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	userId, exists := context.Get("userId")

	if !exists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "user doesn't exist"})
		return
	}

	userIdStr, ok := userId.(string)
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"message": "user ID is not a string"})
		return
	}

	// Convert the string to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		fmt.Println("Error converting string to ObjectID:", err)
		return
	}

	com.Creator = objectID
	com.OnApproved = false
	com.SubsribersCount = 0
	com.DateCreated = primitive.NewDateTimeFromTime(time.Now())

	comm, err1 := com.Save()

	if err1 != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "community requested", "community": comm})

}
