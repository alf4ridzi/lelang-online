package models

import (
	"time"
)

type Auction struct {
	ID     uint `gorm:"primaryKey"`
	ItemID int  `json:"item_id" binding:"required"`
	Item   Item `json:"item"`
	UserID int
	User   User `json:"user"`

	StartTime time.Time `json:"start_time" binding:"required" time_format:"2006-01-02"`
	EndTime   time.Time `json:"end_time" binding:"required" time_format:"2006-01-02"`

	FinalPrice  *int64 `json:"final_price"`
	StartingBid int64  `json:"starting_bid" binding:"required"`

	WinnerID *int
	Winner   *User `json:"winner"`

	CurrentBid int64 `json:"current_bid"`
	BidCount   int   `json:"bid_count"`

	Status    int `gorm:"default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuctionRequest struct {
	ItemID      int       `json:"item_id" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	StartingBid int64     `json:"starting_bid" binding:"required"`
}
