package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AccountRequestForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Signin(ctx *gin.Context, db *sql.DB) {
	var input AccountRequestForm

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// パスワード検証
	var storedPassHash string
	if err := db.QueryRow("SELECT hashed_pass FROM users WHERE name=?", input.Username).Scan(&storedPassHash); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		} else {
			fmt.Printf("getPassHashQuery: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// ハッシュパスワードの照合
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassHash), []byte(input.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}
