package main

import (
	"context"
	"errors"
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
	sourcesService := sourcesUseCases.NewSourceService(sourcesRepo)
	sourcesHandler := sourcesAdapter.NewSourcesHandler(sourcesService)

	handlers := http.NewServeMux()
	handlers.HandleFunc("POST /sources", sourcesHandler.CreateSource)
	handlers.HandleFunc("GET /sources", sourcesHandler.GetAllSources)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handlers,
	}

	go func() {
		ctx := context.Background()
		logger.Info(ctx, "Starting HTTP server")

		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Error(ctx, err.Error())
			os.Exit(int(syscall.SIGINT))
		}
		logger.Info(ctx, "Stop serving new connections")
	}()

	shutdownServer := make(chan os.Signal, 1)
	signal.Notify(shutdownServer, syscall.SIGINT, syscall.SIGTERM)
	<-shutdownServer

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error(ctx, "Server forced to shutdown", err.Error())
	}

	logger.Info(ctx, "Server gracefully stopped")
}
