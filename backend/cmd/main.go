package main

import (
	"lelang-online-api/config"
	"lelang-online-api/database"
	"lelang-online-api/database/migrations"
	"lelang-online-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.InitEnv(); err != nil {
		log.Fatal(err)
	}

	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	err = migrations.Migrate(db, config.ModelMigration)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.SetupRoutes(router)
	router.Run(":8080")
}
