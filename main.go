package main

import (
    "url-shortener/config"
    "url-shortener/routes"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
    config.ConnectDatabase()
    r := routes.SetupRouter()
    r.Static("/docs", "./docs")
    url := ginSwagger.URL("/docs/swagger.json")
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
    r.Run(":8080")
}
