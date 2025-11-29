package models

import "time"

type Item struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	StartingBid int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
