package helpers

import "github.com/gin-gonic/gin"

func ResponseJson(ctx *gin.Context, statusCode int, status bool, message string, data any) {
	ctx.JSON(statusCode, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
