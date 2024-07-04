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

func main() {

	config.ConnectDatabase()

	r := routes.SetupRouter()
	r.Static("/docs", "./docs")
	url := ginSwagger.URL("/docs/swagger.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":8080")
}
