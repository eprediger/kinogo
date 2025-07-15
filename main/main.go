package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	logLevel "application/ports/logging"
	sourcesUseCases "application/use_cases/sources"
	sourcesAdapter "infrastructure/http/source"
	logger "infrastructure/logging"
	ddbb "infrastructure/repositories/ddbb"
)

func main() {
	logger := logger.NewLogger(logLevel.Debug)
	logger.Info(context.Background(), "Starting application")

	sourcesRepo := ddbb.NewSourcesRepository(logger)
	sourcesService := sourcesUseCases.NewSourcesService(sourcesRepo)
	sourcesHandler := sourcesAdapter.NewSourcesHandler(sourcesService)

	handlers := http.NewServeMux()
	handlers.HandleFunc("POST /sources", sourcesHandler.CreateSource)
	handlers.HandleFunc("GET /sources", sourcesHandler.GetAllSources)
	server := &http.Server{
		Addr:    ":8080",
		Handler: handlers,
	}

	go func() {
		logger.Info(context.Background(), "Application started")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown")
	} else {
		log.Println("Server gracefully stopped")
	}
}
