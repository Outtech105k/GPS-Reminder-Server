package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Outtech105k/GPS-Reminder-Server/web/router"
	"github.com/gin-gonic/gin"
)

func main() {
	handler := gin.Default()
	router.SetRoutes(handler)

	srv := &http.Server{
		Addr:    ":80",
		Handler: handler,
	}

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

	log.Println("Server exiting")
}
