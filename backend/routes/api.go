package routes

import (
	"lelang-online-api/controllers"
	"lelang-online-api/repositories"
	"lelang-online-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, r *gin.Engine) {
	{
		api := r.Group("api")
		authRoute(db, api)
	}
}

func authRoute(db *gorm.DB, route *gin.RouterGroup) {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewAuthController(&userService)

	{
		authGroup := route.Group("auth")
		authGroup.POST("/login", userController.Login)
		authGroup.POST("/register", userController.Register)
	}

}
