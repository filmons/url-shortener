package models

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at,omitempty"`
    Name      string    `json:"name"`
    Email     string    `json:"email" gorm:"unique"`
    Password  string    `json:"password"` 
    URLs      []URL     `gorm:"foreignKey:UserID"`
}
