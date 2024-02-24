package main

import (
	"ginapp/config"
	"ginapp/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("error loading the config file")
	}

	db, err := database.ConnectDatabase(cfg)

	if err != nil {

		log.Fatalf("error connecting to the database", err)
	}

	router := gin.Default()
	userGroup := router.Group("/")
	userGroup.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "User route"})
	})

	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("localhost error %v", err)
	}

}
