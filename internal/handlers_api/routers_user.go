package handlers

import (
	"Effective_Mobile/internal/services"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", h.CreateUser).Methods("POST")
	router.HandleFunc("/users", h.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", h.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
}
