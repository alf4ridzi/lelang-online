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

	if user == nil {
		helpers.ResponseJson(ctx, http.StatusNotFound, false, "username atau password salah", nil)
		return
	}

	session := sessions.Default(ctx)
	session.Set("user_id", user.ID)
	session.Save()

	helpers.ResponseJson(ctx, http.StatusOK, true, "berhasil login", nil)
}

func (c *AuthController) Register(ctx *gin.Context) {
	var register models.Register
	if err := ctx.ShouldBindJSON(&register); err != nil {
		helpers.ResponseJson(ctx, http.StatusBadRequest, false, "bad request", nil)
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	if register.Password != register.ConfirmPassword {
		helpers.ResponseJson(ctx, http.StatusBadRequest, false, "password tidak sama", nil)
		return
	}

	user := models.User{
		Name:     register.Name,
		Username: register.Username,
		Password: register.Password,
		Phone:    register.Phone,
		RoleID:   3,
	}

	if err := c.service.Register(reqCtx, user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			helpers.ResponseJson(ctx, http.StatusConflict, false, "username sudah ada", nil)
			return
		}

		helpers.ResponseJson(ctx, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	helpers.ResponseJson(ctx, http.StatusOK, true, "berhasil register", nil)
}
