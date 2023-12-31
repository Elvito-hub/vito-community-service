package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"vitocom.com/community/db"
	"vitocom.com/community/routes"
)

func main() {

	client, err := db.GetMongoClient()

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":5214")
}
