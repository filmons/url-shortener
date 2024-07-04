// package config

// import (
//     "log"
//     "url-shortener/models"

//     "gorm.io/driver/mysql"
//     "gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDatabase() {
//     dsn := "root:dbfilmon@tcp(127.0.0.1:3306)/url_shortener?charset=utf8mb4&parseTime=True&loc=Local"
//     database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//     if err != nil {
//         log.Fatal("Failed to connect to database:", err)
//     }

//     // Migrer automatiquement le schéma de la base de données
//     err = database.AutoMigrate(&models.URL{})
//     if err != nil {
//         log.Fatal("Failed to migrate database schema:", err)
//     }

//     DB = database
// }

package config

import (
	"fmt"
	"log"
	"os"
	"url-shortener/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrer automatiquement le schéma de la base de données
	err = database.AutoMigrate(&models.User{}, &models.URL{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	DB = database
}
