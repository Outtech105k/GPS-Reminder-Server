package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetHandler(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "pong",
		})
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "Not Found",
		})
	})
}
