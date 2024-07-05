package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequestForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var input LoginRequestForm

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: This is TEMP.
	if input.Username == "user" && input.Password == "pass" {
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello, %s!", input.Username),
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid username or password",
		})
	}
}
