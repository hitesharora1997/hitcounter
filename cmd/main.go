package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hitesharora1997/hitcounter/internal/server"
)

func main() {
	serverInstance := server.NewServer("server_data.json", time.Now)
	httpServer := &http.Server{
		Addr:    ":8090",
		Handler: serverInstance,
	}
	go func() {
		log.Printf("Starting server on port: %s", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("HTTP server stopped:", err)
		}
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Received shutdown signal. Initiating graceful shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Println("Error during server shutdown:", err)
	}
	serverInstance.Stop()
	log.Println("Server gracefully stopped.")
}
