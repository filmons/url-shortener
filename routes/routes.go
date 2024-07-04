package routes

import (
    "url-shortener/controllers"
    "url-shortener/middlewares"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "time"
    
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Ajouter le middleware CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

    // Routes publiques
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    r.GET("/urls", controllers.GetUserURLs)
    r.GET("/:short_url", controllers.RedirectURL)
    // Routes protégées
    protected := r.Group("/")
    protected.Use(middlewares.AuthMiddleware())
    {
        protected.POST("/shorten", controllers.CreateShortURL)
        // protected.GET("/urls", controllers.GetUserURLs)
        // protected.GET("/:short_url", controllers.RedirectURL)
    }

    return r
}
// package routes

// import (
//     "github.com/gin-gonic/gin"
//     "github.com/gin-contrib/cors"
//     "time"
//     "url-shortener/controllers"
//     "url-shortener/middlewares"
// )

// func SetupRouter() *gin.Engine {
//     r := gin.Default()

//     // Ajouter le middleware CORS
//     r.Use(cors.New(cors.Config{
//         AllowOrigins:     []string{"http://localhost:3000"},
//         AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
//         AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
//         ExposeHeaders:    []string{"Content-Length"},
//         AllowCredentials: true,
//         MaxAge: 12 * time.Hour,
//     }))

//     // Routes publiques
//     r.POST("/register", controllers.Register)
//     r.POST("/login", controllers.Login)

//     // Routes de redirection (publique)
//     r.GET("/:short_url", controllers.RedirectURL)
//     r.POST("/shorten", controllers.CreateShortURL)

//     // Routes protégées
//     protected := r.Group("/")
//     protected.Use(middlewares.AuthMiddleware())
//     {
//         // protected.POST("/shorten", controllers.CreateShortURL)
//         protected.GET("/urls", controllers.GetUserURLs)
//     }

//     return r
// }
