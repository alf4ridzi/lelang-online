package models

import (
	"time"
)

type Auction struct {
	ID     uint `gorm:"primaryKey"`
	ItemID int
	Item   Item
	UserID int
	User   User

	StartTime time.Time
	EndTime   time.Time

	FinalPrice  int64
	StartingBid int64

	WinnerID *int
	Winner   *User

	CurrentBid int64
	BidCount   int

	Status    int `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
