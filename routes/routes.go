package routes

import (
	"github.com/gin-gonic/gin"
	"vitocom.com/community/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/allcommunities", getAllCommunities)
	server.POST("/createcommunity", middlewares.Authenticate, createCommunity)

}
