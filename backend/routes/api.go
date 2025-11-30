package routes

import (
	"lelang-online-api/controllers"
	"lelang-online-api/middlewares"
	"lelang-online-api/repositories"
	"lelang-online-api/services"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, r *gin.Engine) {
	r.Use(middlewares.Session(os.Getenv("SECRET_KEY")))

	{
		api := r.Group("api")
		authRoute(db, api)

		api.Use(middlewares.AuthMiddleware)
		userRoute(db, api)
	}
}

func userRoute(db *gorm.DB, route *gin.RouterGroup) {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	{
		usersGroup := route.Group("users")
		usersGroup.GET("/profile", userController.Profile)
	}
}

func authRoute(db *gorm.DB, route *gin.RouterGroup) {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewAuthController(userService)

	{
		authGroup := route.Group("auth")
		authGroup.POST("/login", userController.Login)
		authGroup.POST("/register", userController.Register)
	}

}
