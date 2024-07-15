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

	setRemindersGroup(router, db, jwtMiddleware)

	// アカウント名・パスワードを入力してトークン取得
	router.POST("/auth/token", jwtMiddleware.LoginHandler)
	router.GET("/auth/token/reflesh", jwtMiddleware.RefreshHandler)

	// 404レスポンス
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})
	})
}

func setRemindersGroup(router *gin.Engine, db *sql.DB, jwtMiddleware *jwt.GinJWTMiddleware) {
	group := router.Group("/reminders")
	group.Use(jwtMiddleware.MiddlewareFunc())

	group.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Reminders",
			"username": auth.GetUsernameInJWT(ctx),
		})
	})

	group.POST("/", func(ctx *gin.Context) {
		handler.PostRemind(ctx, db)
	})
}
