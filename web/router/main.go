package router

import (
	"database/sql"
	"net/http"

	"github.com/Outtech105k/GPS-Reminder-Server/web/handler"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, db *sql.DB) {
	router.POST("/signup", func(ctx *gin.Context) {
		handler.Signup(ctx, db)
	})

	router.POST("/signin", func(ctx *gin.Context) {
		handler.Signin(ctx, db)
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "Not Found",
		})
	})
}
