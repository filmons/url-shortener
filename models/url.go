package models

import "time"

type URL struct {
    ID            uint      `gorm:"primaryKey" json:"id"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
    DeletedAt     *time.Time `json:"deleted_at,omitempty"`
    LongURL       string    `json:"long_url"`
    ShortURL      string    `json:"short_url" gorm:"unique"`
    UserID        uint      `json:"user_id"`
}
