package controllers

import (
	"context"
	"lelang-online-api/helpers"
	"lelang-online-api/models"
	"lelang-online-api/services"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuctionController struct {
	service services.AuctionService
}

func NewAuctionController(service services.AuctionService) *AuctionController {
	return &AuctionController{service: service}
}

func (c *AuctionController) Bid(ctx *gin.Context) {
	var req models.BidRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.ResponseJson(ctx, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	session := sessions.Default(ctx)
	val := session.Get("user_id")
	userID := val.(uint)

	if err := c.service.AddBid(reqCtx, userID, &req); err != nil {
		helpers.ResponseJson(ctx, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	helpers.ResponseJson(ctx, http.StatusOK, true, "berhasil membuat bid", nil)

}

func (c *AuctionController) New(ctx *gin.Context) {
	var req models.AuctionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {

	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	session := sessions.Default(ctx)
	userIDStr := session.Get("user_id")
	userID := userIDStr.(uint)

	auction := models.Auction{
		UserID:      userID,
		ItemID:      req.ItemID,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		StartingBid: req.StartingBid,
		Status:      1,
	}

	if err := c.service.NewAuction(reqCtx, &auction); err != nil {
		helpers.ResponseJson(ctx, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	helpers.ResponseJson(ctx, http.StatusOK, true, "berhasil menambah lelang", nil)
}

func (c *AuctionController) All(ctx *gin.Context) {
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	auctions, err := c.service.All(reqCtx)
	if err != nil {
		helpers.ResponseJson(ctx, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	data := map[string][]models.Auction{
		"auctions": auctions,
	}

	helpers.ResponseJson(ctx, http.StatusOK, true, "berhasil", data)
}
