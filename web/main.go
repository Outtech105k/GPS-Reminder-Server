package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Outtech105k/GPS-Reminder-Server/web/auth"
	"github.com/Outtech105k/GPS-Reminder-Server/web/db"
	"github.com/Outtech105k/GPS-Reminder-Server/web/router"
	"github.com/gin-gonic/gin"
)

// @title           GPS Reminder API
// @version         Beta
// @description     位置情報と連携したリマインダーアプリに使用するAPIです

// @license.name  MIT License
// @license.url   https://opensource.org/license/mit
func main() {
	fmt.Println("--- Server starting ---")

	// DB接続確立
	db, err := db.Connect()
	if err != nil {
		fmt.Printf("DB: %v\n", err)
	}
	defer db.Close()

	// JWT認証設定
	jwtMiddleWare, err := auth.NewJWTMiddleware(db)
	if err != nil {
		fmt.Printf("JWT: %v\n", err)
	}

	// Ginサーバセットアップ
	handler := gin.Default()
	router.SetRoutes(handler, db, jwtMiddleWare)

	srv := &http.Server{
		Addr:    ":80",
		Handler: handler,
	}

	// サーバ起動
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Panic("Server forced to shutdown:", err)
	}

	fmt.Println("--- Server exiting ---")
}
