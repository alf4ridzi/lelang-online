package controllers

import (
	"context"
	"lelang-online-api/helpers"
	"lelang-online-api/services"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (c *UserController) Profile(ctx *gin.Context) {

	session := sessions.Default(ctx)
	val := session.Get("user_id")

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	user, err := c.service.Profile(reqCtx, val)
	if err != nil {
		helpers.ResponseJson(ctx, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	user.Password = ""

	helpers.ResponseJson(ctx, http.StatusOK, true, "ok", user)
}

func (c *UserController) GetItems(ctx *gin.Context) {

}
