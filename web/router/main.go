package router

import (
	"net/http"

	"github.com/Outtech105k/GPS-Reminder-Server/web/handler"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "pong",
		})
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "Not Found",
		})
	})

	router.POST("/login", handler.Login)
}
