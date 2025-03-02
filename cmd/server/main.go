package main

import (
	"log/slog"
	"net/http"
	"os"
	"test-project/internal/config"
	repository "test-project/internal/database/postgres"
)

func main() {
	if err := run(); err != nil {
		log := setupLogger("local")
		log.Error("application terminated with an error")
		os.Exit(1)
	}
}

func run() error {
	cfg := config.Load()

	log := setupLogger(cfg.Env)

	log.Info(
		"starting book service",
		slog.String("env", cfg.Env),
		slog.String("version", "123"),
	)
	log.Debug("debug messages are enabled")

	_, err := repository.Connect(cfg)
	if err != nil {
		log.Error("database connection failed", slog.Any("error", err))
		os.Exit(1)
	}

	log.Info("Successfull connect to the database")
	log.Info("Successful connect to the database")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	log.Info("Server is running on port 8080")
	return http.ListenAndServe(":8080", nil)
}

func setupLogger(env string) *slog.Logger {
	var level slog.Level

	switch env {
	case "local":
		level = slog.LevelDebug
	case "production":
		level = slog.LevelInfo
	default:
		level = slog.LevelWarn
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	return slog.New(handler)
}
