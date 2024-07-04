package main

import (
	"Effective_Mobile/internal/config"
	initApp "Effective_Mobile/internal/initApp"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "Effective_Mobile/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Effective_Mobile API
// @version 1.0
// @description This is the API for the Effective_Mobile application.
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("config", cfg))

	application, err := initApp.InitializeApp(cfg, log)
	if err != nil {
		log.Error("failed to initialize application", slog.String("error", err.Error()))
		os.Exit(1)
	}

	http.Handle("/swagger/", config.Cors(httpSwagger.WrapHandler))
	go func() {
		log.Info("starting Swagger documentation server")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Error("failed to serve swagger", slog.String("error", err.Error()))
		}
	}()

	go application.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	log.Info("application stopped", slog.String("signal", sign.String()))
	application.Stop()
	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "dev":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
