package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"

	"github.com/Outtech105k/GPS-Reminder-Server/web/auth"
	"github.com/Outtech105k/GPS-Reminder-Server/web/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PostUsers(ctx *gin.Context, db *sql.DB) {
	var input auth.AccountRequest

	// リクエスト不備チェック
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorDefaultResponse{
			Error: "Invalid request payload",
		})
		return
	}

	// ユーザー名バリデーション
	if isInvalidUsername(input.Username) {
		ctx.JSON(http.StatusBadRequest, response.ErrorDefaultResponse{
			Error: "Username did not satisfy the requirements",
		})
		return
	}

	// パスワードバリデーション
	if !isValidPassword(input.Password) {
		ctx.JSON(http.StatusBadRequest, response.ErrorDefaultResponse{
			Error: "Password did not satisfy the requirements",
		})
		return
	}

	// ユーザ名の重複(既存)チェック
	var existingUser string
	err := db.QueryRow("SELECT name FROM users WHERE name=? LIMIT 1", input.Username).Scan(&existingUser)
	// err == nil           の時、レコードが存在するので、重複扱い (離脱)
	// err == sql.ErrNoRows の時、レコードが存在しないので、正常   (続行)
	// それ以外の時、Queryエラー                                   (離脱)
	if err == nil {
		ctx.JSON(http.StatusConflict, response.ErrorDefaultResponse{Error: "Username already exists"})
		return
	} else if err != sql.ErrNoRows {
		fmt.Printf("usernameConflictQuery: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, response.ErrorDefaultResponse{Error: "Database error"})
		return
	}

	// パスワードのハッシュ化
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("generatePassword: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, response.ErrorDefaultResponse{
			Error: "User resist error",
		})
		return
	}

	// ユーザ登録実行
	_, err = db.Exec("INSERT INTO users(name, hashed_pass) VALUES(?, ?)", input.Username, string(hashBytes))
	if err != nil {
		fmt.Printf("registerUser: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, response.ErrorDefaultResponse{
			Error: "User resist error",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessDefaultResponse{
		Message: "User registerd successfully",
	})
}

func isInvalidUsername(name string) bool {
	return !regexp.MustCompile(`^\w{3,24}$`).MatchString(name)
}

func isValidPassword(password string) bool {
	var (
		hasMinLen   = len(password) >= 8 && len(password) <= 24
		hasUpper    = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower    = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber   = regexp.MustCompile(`\d`).MatchString(password)
		hasSpecial  = regexp.MustCompile(`[!@#$%^&*]`).MatchString(password)
		hasNoSpaces = !regexp.MustCompile(`\s`).MatchString(password)
	)
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial && hasNoSpaces
}
