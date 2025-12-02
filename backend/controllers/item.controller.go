package controllers

import (
	"context"
	"errors"
	"lelang-online-api/helpers"
	"lelang-online-api/models"
	"lelang-online-api/services"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemController struct {
	service services.ItemService
}

func NewItemController(service services.ItemService) *ItemController {
	return &ItemController{service: service}
}
func (c *ItemController) GetByID(ctx *gin.Context) {
	userID := ctx.Param("id")

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	items, err := c.service.GetByID(reqCtx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.ResponseJson(ctx, http.StatusNotFound, false, "no items", nil)
			return
		}

		helpers.ResponseJson(ctx, http.StatusInternalServerError, false, "internal server error", nil)
		return
	}

	data := map[string][]models.Item{
		"items": items,
	}

	helpers.ResponseJson(ctx, http.StatusOK, true, "berhasil", data)
}

func (c *ItemController) Store(ctx *gin.Context) {
	var itemReq models.Item
	if err := ctx.ShouldBindJSON(&itemReq); err != nil {
		helpers.ResponseJson(ctx, http.StatusBadRequest, false, "bad request", nil)
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	session := sessions.Default(ctx)
	val := session.Get("user_id")
	userID, ok := val.(uint)

	if !ok {
		helpers.ResponseJson(ctx, http.StatusInternalServerError, false, "internal server error", nil)
		return
	}

	item := models.Item{
		UserID:      userID,
		Name:        itemReq.Name,
		Description: itemReq.Description,
	}

	if err := c.service.Create(reqCtx, item); err != nil {
		helpers.ResponseJson(ctx, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	helpers.ResponseJson(ctx, http.StatusOK, true, "berhasil menyimpan barang baru", nil)
}
