package router

import (
	"database/sql"
	"net/http"

	"github.com/Outtech105k/GPS-Reminder-Server/web/auth"
	"github.com/Outtech105k/GPS-Reminder-Server/web/handler"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, db *sql.DB, jwtMiddleware *jwt.GinJWTMiddleware) {
	// ユーザ登録
	router.POST("/users", func(ctx *gin.Context) {
		handler.PostUsers(ctx, db)
	})

	// アカウント名・パスワードを入力してトークン取得
	router.POST("/auth/token", jwtMiddleware.LoginHandler)
	router.GET("/auth/token/reflesh", jwtMiddleware.RefreshHandler)

	// トークン認証
	authGroup := router.Group("/reminders")
	authGroup.Use(jwtMiddleware.MiddlewareFunc())
	authGroup.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Reminders",
			"username": auth.GetUsernameInJWT(ctx),
		})
	})

	// 404レスポンス
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})
	})
}
