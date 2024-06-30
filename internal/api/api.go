package api

import (
	"Effective_Mobile/internal/repositories"
	"Effective_Mobile/internal/services"
	handlers "Effective_Mobile/internal/handlers_api"
	"database/sql"

	"github.com/gorilla/mux"
)

type API struct {
	DB     *sql.DB
	Router *mux.Router
}

func NewAPI(router *mux.Router, dbConn *sql.DB) (*API, error) {
	return &API{
		DB:     dbConn,
		Router: router,
	}, nil
}

//TODO: надо разделить userHandler и worklogHandler 
func (a *API) RegisterHandlers() {
	userRepo := repositories.NewUserRepository(a.DB)
	worklogRepo := repositories.NewWorklogRepository(a.DB)

	userService := services.NewUserService(userRepo)
	worklogService := services.NewWorklogService(worklogRepo)

	userHandler := handlers.NewTaskHandler(userService)
	worklogHandler := handlers.NewTaskHandler(worklogService)

	userHandler.RegisterRoutes(a.Router)
	worklogHandler.RegisterRoutes(a.Router)
}