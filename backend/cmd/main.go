package main

import (
	"lelang-online-api/config"
	"lelang-online-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.InitEnv(); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.SetupRoutes(router)
	router.Run(":8080")
}
