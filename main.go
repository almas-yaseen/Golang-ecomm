package main

import (
	"ginapp/config"
	"ginapp/database"
	routes "ginapp/router"

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

		log.Fatalf("error connecting to the database %v", err)
	}

	router := gin.Default()
	routes.UserRoutes(router.Group("/"), db)
	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("localhost error %v", err)
	}

}
