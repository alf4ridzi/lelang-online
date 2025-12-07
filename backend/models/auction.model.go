package models

import (
	"time"
)

type Auction struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	ItemID uint `json:"item_id" binding:"required"`
	Item   Item `json:"item" gorm:"foreignKey:ItemID;references:ID;constraint:OnDelete:CASCADE"`
	UserID uint
	User   User `json:"user"`

	StartTime time.Time `json:"start_time" binding:"required" time_format:"2006-01-02"`
	EndTime   time.Time `json:"end_time" binding:"required" time_format:"2006-01-02"`

	FinalPrice  *int64 `json:"final_price"`
	StartingBid int64  `json:"starting_bid" binding:"required"`

	WinnerID *uint
	Winner   *User `json:"winner"`

	CurrentBid int64 `json:"current_bid"`
	BidCount   int   `json:"bid_count"`

	Status int `gorm:"default:1"`

	Histories []AuctionHistory `gorm:"foreignKey:AuctionID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuctionRequest struct {
	ItemID      uint      `json:"item_id" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	StartingBid int64     `json:"starting_bid" binding:"required"`
}
