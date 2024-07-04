// package middlewares

// import (
//     "log"
//     "net/http"
//     "strings"

//     "github.com/dgrijalva/jwt-go"
//     "github.com/gin-gonic/gin"
// )

// var jwtKey = []byte("my_secret_key")

// type Claims struct {
//     UserID uint `json:"userID"`
//     jwt.StandardClaims
// }

// func AuthMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         tokenString := c.GetHeader("Authorization")
//         if tokenString == "" {
//             log.Println("No access token provided")
//             c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
//             c.Abort()
//             return
//         }

//         log.Println("Token String:", tokenString)

//         claims := &Claims{}
//         token, err := jwt.ParseWithClaims(strings.TrimSpace(strings.Replace(tokenString, "Bearer ", "", 1)), claims, func(token *jwt.Token) (interface{}, error) {
//             return jwtKey, nil
//         })

//         if err != nil {
//             log.Println("Error parsing token:", err)
//             c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
//             c.Abort()
//             return
//         }

//         if !token.Valid {
//             log.Println("Token is invalid")
//             c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
//             c.Abort()
//             return
//         }

//         log.Println("Token is valid. UserID:", claims.UserID)
//         c.Set("userID", claims.UserID)
//         c.Next()
//     }
// }
package middlewares

import (
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
            c.Abort()
            return
        }

        claims := &Claims{}
        token, err := jwt.ParseWithClaims(strings.TrimSpace(strings.Replace(tokenString, "Bearer ", "", 1)), claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        if !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("userID", claims.UserID)
        c.Next()
    }
}
