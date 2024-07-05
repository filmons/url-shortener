// Description: This file contains the URL model which is used to store the URL data in the database.
package models

import (
	"gorm.io/gorm"
	"time"
)

type URL struct {
	gorm.Model
	LongURL        string     `json:"long_url"`
	ShortURL       string     `json:"short_url" gorm:"unique"`
	UserID         uint       `json:"user_id"`
	Clicks         int        `json:"clicks" gorm:"default:0"`
	LastAccessedAt *time.Time `json:"last_accessed_at"`
}
