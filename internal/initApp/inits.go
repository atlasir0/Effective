package init

import (
	"Effective_Mobile/internal/app"
	"Effective_Mobile/internal/config"
	"Effective_Mobile/internal/repositories"
	"Effective_Mobile/internal/services"
	"Effective_Mobile/internal/storage/postgres"
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeApp(cfg *config.Config, log *slog.Logger) (*app.App, error) {
	db, _, err := postgres.InitDB()
	if err != nil {
		return nil, err
	}
	defer postgres.CloseDB(db)

	connString := cfg.Database.ConnectionString()

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	dbPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	userRepo, err := repositories.NewUserRepository(dbPool)
	if err != nil {
		return nil, err
	}
	worklogRepo, err := repositories.NewWorklogRepository(dbPool)
	if err != nil {
		return nil, err
	}

	userService := services.NewUserService(userRepo)
	worklogService := services.NewWorklogService(worklogRepo)

	application := app.New(log, cfg.HTTP.Port, userService, worklogService)
	return application, nil
}
