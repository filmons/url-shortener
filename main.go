package main

import (

	"log"
	"url-shortener/config"
	"url-shortener/routes"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	// Find .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}


// @title URL Shortener API
// @version 1.0
// @description This is a sample server for a URL shortener.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
    config.ConnectDatabase()

    r := routes.SetupRouter()
    r.Static("/docs", "./docs")
    url := ginSwagger.URL("/docs/swagger.json") // The url pointing to API definition
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    r.Run(":8080")
}
