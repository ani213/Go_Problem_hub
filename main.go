package main

import (
	"os"

	"github.com/ani213/Problemhub_backend/config"
	"github.com/ani213/Problemhub_backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	router := gin.Default()
	routes.Routes(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
