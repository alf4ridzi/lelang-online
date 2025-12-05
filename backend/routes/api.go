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
		authRoutes(db, api)

		api.Use(middlewares.AuthMiddleware)
		userRoutes(db, api)
		ItemRoutes(db, api)
		AuctionRoutes(db, api)
	}
}

func userRoutes(db *gorm.DB, route *gin.RouterGroup) {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	{
		usersGroup := route.Group("users")
		usersGroup.GET("/profile", userController.Profile)
		usersGroup.GET("/items", userController.GetItems)
	}
}

func authRoutes(db *gorm.DB, route *gin.RouterGroup) {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewAuthController(userService)

	{
		authGroup := route.Group("auth")
		authGroup.POST("/login", userController.Login)
		authGroup.POST("/register", userController.Register)
		authGroup.DELETE("/logout", userController.Logout)
	}
}

func ItemRoutes(db *gorm.DB, route *gin.RouterGroup) {
	itemRepo := repositories.NewItemRepo(db)
	itemService := services.NewItemService(itemRepo)
	itemController := controllers.NewItemController(itemService)

	{
		itemGroup := route.Group("items")
		itemGroup.POST("", itemController.Store)
		itemGroup.GET("/:id", itemController.GetByID)
		itemGroup.PUT("/:id", itemController.Update)
		itemGroup.DELETE("/:id", itemController.Delete)
	}
}

func AuctionRoutes(db *gorm.DB, route *gin.RouterGroup) {
	auctionRepo := repositories.NewAuctionRepo(db)
	auctionService := services.NewAuctionService(auctionRepo)
	auctionController := controllers.NewAuctionController(auctionService)

	{
		auctionGroup := route.Group("auctions")
		auctionGroup.POST("/new", auctionController.New)
		auctionGroup.GET("", auctionController.All)
	}
}
