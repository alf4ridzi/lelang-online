package controllers

import (
	"context"
	"lelang-online-api/helpers"
	"lelang-online-api/models"
	"lelang-online-api/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service services.UserService
}

func NewAuthController(service services.UserService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var login models.Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		helpers.ResponseJson(ctx, http.StatusBadRequest, false, "bad request", nil)
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	user, err := c.service.Login(reqCtx, login)
	if err != nil {
		helpers.ResponseJson(ctx, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	helpers.ResponseJson(ctx, http.StatusOK, true, "berhasil login", user)
}

func (c *AuthController) Register(ctx *gin.Context) {

}
