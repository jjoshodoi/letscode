package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"letscode/project-01-url-shortener/internal/handlers"
	"letscode/project-01-url-shortener/internal/logging"
	"letscode/project-01-url-shortener/internal/store"
)

type Config struct {
	Port            string
	StorePath       string
	LogLevel        string
	ShutdownTimeout time.Duration
}

func loadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else if port[0] != ':' {
		port = ":" + port
	}
	storePath := os.Getenv("STORE_PATH")
	if storePath == "" {
		storePath = "data/urls.json"
	}
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	return Config{
		Port:            port,
		StorePath:       storePath,
		LogLevel:        logLevel,
		ShutdownTimeout: 5 * time.Second,
	}
}

func main() {
	cfg := loadConfig()
	logging.Init(cfg.LogLevel)

	st := store.NewFileStore(cfg.StorePath)
	h := handlers.NewHandler(st)

	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: h.Routes(),
	}

	logging.Infof("starting server on %s", cfg.Port)

	// start server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Fatalf("server failed: %v", err)
		}
	}()

	// wait for interrupt
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	logging.Infof("shutdown signal received, shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logging.Errorf("server shutdown failed: %v", err)
	} else {
		logging.Infof("server stopped gracefully")
	}
}
