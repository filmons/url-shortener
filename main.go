package main

import (
    "url-shortener/config"
    "url-shortener/routes"
     ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
)

func main() {
    
    config.ConnectDatabase()

    r := routes.SetupRouter()
    r.Static("/docs", "./docs")
    url := ginSwagger.URL("/docs/swagger.json") // The url pointing to API definition
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    r.Run(":8080")
}
