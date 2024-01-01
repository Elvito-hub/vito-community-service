package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vitocom.com/community/utils"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
	}
	userClaims, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
	}

	context.Set("userId", userClaims.UserId)
	context.Set("role", userClaims.Role)

}
