package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	{
		api := r.Group("api")
		authRoute(api)
	}
}

func authRoute(route *gin.RouterGroup) {
	route.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": "hello world",
		})
	})
}
