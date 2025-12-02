package models

import "time"

type Item struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	UserID      uint
	User        User
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
