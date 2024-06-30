package api

import (
	"Effective_Mobile/internal/services"
	"github.com/gorilla/mux"

)

type UserHandler struct {
	UserService *services.UserService
}

func NewTaskHandler(taskService *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: taskService,
	}
}
//TODO: разделить на worklogs и users
func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/users", GetUsersHandler).Methods("GET")
	router.HandleFunc("/users", CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/paginated", GetPaginatedUsersHandler).Methods("GET")
	router.HandleFunc("/users/filter", GetFilteredUsersHandler).Methods("GET")
	router.HandleFunc("/users/{userId}", GetUserByIDHandler).Methods("GET")
	router.HandleFunc("/users/{userId}", UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{userId}", DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/users/{userId}/worklogs", GetUserWorklogsHandler).Methods("GET")
	router.HandleFunc("/users/{userId}/tasks/{taskId}/start", StartTaskHandler).Methods("POST")
	router.HandleFunc("/users/{userId}/tasks/{taskId}/stop", StopTaskHandler).Methods("POST")

}

