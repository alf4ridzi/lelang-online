package main

import (
	"flag"
	"lelang-online-api/config"
	"lelang-online-api/database"
	"lelang-online-api/database/migrations"
	"lelang-online-api/database/seeders"
	"lelang-online-api/handlers"
	"lelang-online-api/middlewares"
	"lelang-online-api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Flag(db *gorm.DB) {
	var seed = flag.Bool("seed", false, "run seeder")
	flag.Parse()

	if *seed {
		seeders.RunSeeder(db)
		os.Exit(0)
	}
}

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

	Flag(db)

	go handlers.HubInstance.Run()

	router := gin.Default()
	router.Use(middlewares.Session(os.Getenv("SECRET_KEY")))
	routes.SetupRoutes(db, router)
	router.Run(":8080")
}
