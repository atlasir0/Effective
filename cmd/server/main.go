package main

import (
	"Effective_Mobile/internal/app"
	"Effective_Mobile/internal/config"
	"Effective_Mobile/internal/repositories"
	"Effective_Mobile/internal/services"
	"Effective_Mobile/internal/storage/postgres"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

//TODO: удалить task и оставить worklogs и users так же сделать пагинацию и норм таймер
//TODO: рефакторинг 
//TODO: сделать нормальное отключение миграций 
//TODO: сделать свагер и покрыть тестами

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("config", cfg))

	log.Debug("debug message")
	db, dbConfig, err := postgres.InitDB()
	if err != nil {
		log.Error("failed to initialize database", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer postgres.CloseDB(db, dbConfig)

	userRepo := repositories.NewUserRepository(db)
	worklogRepo := repositories.NewWorklogRepository(db)

	userService := services.NewUserService(userRepo)
	worklogService := services.NewWorklogService(worklogRepo)

	application := app.New(log, cfg.HTTP.Port, userService, worklogService)

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
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
