package api

import (
	handlers "Effective_Mobile/internal/handlers_api"
	"Effective_Mobile/internal/repositories"
	"Effective_Mobile/internal/services"
	"database/sql"

	"github.com/gorilla/mux"
)

type API struct {
	DB     *sql.DB
	Router *mux.Router
}

func NewAPI(router *mux.Router, dbConn *sql.DB) *API {
	return &API{
		DB:     dbConn,
		Router: router,
	}
}

func (a *API) RegisterHandlers() {
	userRepo := repositories.NewUserRepository(a.DB)
	worklogRepo := repositories.NewWorklogRepository(a.DB)

	userService := services.NewUserService(userRepo)
	worklogService := services.NewWorklogService(worklogRepo)

	userHandler := handlers.NewUserHandler(userService)
	worklogHandler := handlers.NewWorklogHandler(worklogService)

	userHandler.RegisterRoutes(a.Router)
	worklogHandler.RegisterRoutes(a.Router)
}
