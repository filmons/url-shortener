package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"url-shortener/controllers"
	"url-shortener/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/:short_url", controllers.RedirectURL)

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/shorten", controllers.CreateShortURL)
		protected.GET("/urls", controllers.GetUserURLs)
		protected.DELETE("/del/:id", controllers.DeleteURL)
		protected.GET("/stats/:id", controllers.GetURLStats)
	}

	return r
}
