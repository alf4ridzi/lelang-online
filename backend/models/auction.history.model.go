package models

import (
	"time"
)

type AuctionHistory struct {
	ID        uint `gorm:"primaryKey"`
	AuctionID int
	Auction   Auction
	ItemID    int
	Item      Item
	UserID    int
	User      User
	Amount    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
