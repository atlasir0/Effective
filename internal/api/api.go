package api

import (
	handlers "Effective_Mobile/internal/handlers_api"
	"Effective_Mobile/internal/repositories"
	"Effective_Mobile/internal/services"
	"log/slog"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type API struct {
	DB     *pgxpool.Pool
	Router *mux.Router
}

func NewAPI(router *mux.Router, dbConn *pgxpool.Pool) *API {
	return &API{
		DB:     dbConn,
		Router: router,
	}
}

func (a *API) RegisterHandlers() {
	userRepo, err := repositories.NewUserRepository(a.DB)
	if err != nil {
		slog.Error("failed to create user repository", slog.String("error", err.Error()))
	}
	worklogRepo, err := repositories.NewWorklogRepository(a.DB)
	if err != nil {
		slog.Error("failed to create worklog repository", slog.String("error", err.Error()))
	}

	userService := services.NewUserService(userRepo)
	worklogService := services.NewWorklogService(worklogRepo)

	userHandler := handlers.NewUserHandler(userService)
	worklogHandler := handlers.NewWorklogHandler(worklogService)

	userHandler.RegisterRoutes(a.Router)
	worklogHandler.RegisterRoutes(a.Router)
}
