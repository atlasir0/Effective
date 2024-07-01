package handlers

import (
	"Effective_Mobile/internal/services"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, userService *services.UserService, worklogService *services.WorklogService) {
	userHandler := NewUserHandler(userService)
	worklogHandler := NewWorklogHandler(worklogService)

	userHandler.RegisterRoutes(router)
	worklogHandler.RegisterRoutes(router)
}

type WorklogHandler struct {
	WorklogService *services.WorklogService
}

func NewWorklogHandler(worklogService *services.WorklogService) *WorklogHandler {
	return &WorklogHandler{
		WorklogService: worklogService,
	}
}

func (h *WorklogHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/worklogs", h.StartTask).Methods("POST")
	router.HandleFunc("/worklogs/{id}", h.StopTask).Methods("PUT")
	router.HandleFunc("/worklogs/user/{id}", h.GetUserWorklogs).Methods("GET")
}
