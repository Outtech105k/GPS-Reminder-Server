package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/Outtech105k/GPS-Reminder-Server/web/response"
	"github.com/gin-gonic/gin"
)

type RemindRequest struct {
	UserName    string `json:"user_name" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Deadline    string `json:"deadline" binding:"required"`
}

func PostRemind(ctx *gin.Context, db *sql.DB) {
	// リクエスト読み込み
	var request RemindRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorDefaultResponse{
			Error: "Invalid request payload",
		})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("beginTransaction: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, response.ErrorDefaultResponse{
			Error: "Database error",
		})
		return
	}

	// ユーザー名読み出し(クエリ正常で続行)
	var userId int
	err = tx.QueryRow("SELECT id FROM users WHERE name=? LIMIT 1", request.UserName).Scan(&userId)
	if err != nil {
		fmt.Printf("searchUsername: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, response.ErrorDefaultResponse{Error: "Database error"})
		tx.Rollback()
		return
	}

	// 時刻変換
	deadline, err := time.Parse(time.RFC3339, request.Deadline)
	if err != nil {
		fmt.Printf("convertToTime: %v\n", err)
		ctx.JSON(http.StatusBadRequest, response.ErrorDefaultResponse{Error: "Invalid time form"})
		tx.Rollback()
		return
	}

	_, err = tx.Exec(
		"INSERT INTO reminders(user_id, name, description, deadline) VALUES(?, ?, ?, ?)",
		userId, request.Name, request.Description, deadline,
	)
	if err != nil {
		fmt.Printf("registerRemind: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, response.ErrorDefaultResponse{
			Error: "Remind resist error",
		})
		tx.Rollback()
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, response.SuccessDefaultResponse{
		Message: "remind registerd successfully",
	})
}
