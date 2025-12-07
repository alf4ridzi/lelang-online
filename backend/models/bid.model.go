package models

type BidRequest struct {
	ID     uint  `json:"auction_id" binding:"required"`
	Amount int64 `json:"amount" binding:"required"`
}
