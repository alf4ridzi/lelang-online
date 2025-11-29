package controllers

import (
	"lelang-online-api/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.UserService
}

func NewAuthController(service *services.UserService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {

}

func (c *AuthController) Register(ctx *gin.Context) {

}
