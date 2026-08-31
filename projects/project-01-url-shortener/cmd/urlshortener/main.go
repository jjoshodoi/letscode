package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"letscode/project-01-url-shortener/internal/handlers"
	"letscode/project-01-url-shortener/internal/logging"
	"letscode/project-01-url-shortener/internal/store"
)

type Config struct {
	Port            string
	StoreType       string
	StorePath       string
	StoreDSN        string
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

	storeType := strings.ToLower(strings.TrimSpace(os.Getenv("STORE_TYPE")))
	if storeType == "" {
		storeType = "sqlite"
	}

	storePath := os.Getenv("STORE_PATH")
	if storePath == "" {
		storePath = "data/urls.json"
	}

	storeDSN := os.Getenv("STORE_DSN")
	if storeDSN == "" {
		if storeType == "sqlite" || storeType == "sqlite3" {
			storeDSN = "data/urls.db"
		} else {
			storeDSN = storePath
		}
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	return Config{
		Port:            port,
		StoreType:       storeType,
		StorePath:       storePath,
		StoreDSN:        storeDSN,
		LogLevel:        logLevel,
		ShutdownTimeout: 5 * time.Second,
	}
}

func main() {
	cfg := loadConfig()
	logging.Init(cfg.LogLevel)

	storeLocation := cfg.StorePath
	if cfg.StoreType == "sqlite" || cfg.StoreType == "sqlite3" {
		storeLocation = cfg.StoreDSN
	}
	st, err := store.NewRepository(cfg.StoreType, storeLocation)
	if err != nil {
		logging.Fatalf("failed to initialize store: %v", err)
	}
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
