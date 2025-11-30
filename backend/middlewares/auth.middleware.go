package middlewares

import (
	"lelang-online-api/helpers"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	if userID == nil {
		helpers.ResponseJson(ctx, http.StatusUnauthorized, false, "unauthorized", nil)
		ctx.Abort()
		return
	}

	ctx.Next()
}
