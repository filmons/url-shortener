package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"` // Assurez-vous que le mot de passe est inclus ici
    URLs     []URL  `gorm:"foreignKey:UserID"`
}
