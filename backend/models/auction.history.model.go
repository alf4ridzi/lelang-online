package models

import (
	"time"
)

type AuctionHistory struct {
	ID        uint `gorm:"primaryKey"`
	AuctionID uint
	Auction   Auction `gorm:"foreignKey:AuctionID;references:ID;constraint:OnDelete:CASCADE"`
	ItemID    uint
	Item      Item `gorm:"foreignKey:ItemID;references:ID;constraint:OnDelete:CASCADE"`
	UserID    uint
	User      User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Amount    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
