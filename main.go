package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"zonghai-api/config"
	"zonghai-api/migrations"
	"zonghai-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	config.InitDB()
	defer config.CloseDB()

	migrations.Migrate()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")

	r := gin.Default()
	r.Use(cors.New(corsConfig))

	routes.Serve(r)
	r.Run(":" + os.Getenv("PORT"))
}
