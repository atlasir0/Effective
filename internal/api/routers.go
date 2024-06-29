package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	
	router.HandleFunc("/users/paginated", GetPaginatedUsersHandler).Methods("GET")
	router.HandleFunc("/users/filter", GetFilteredUsersHandler).Methods("GET")
	router.HandleFunc("/users/{userId}", GetUserByIDHandler).Methods("GET")
	router.HandleFunc("/users", GetUsersHandler).Methods("GET")
	router.HandleFunc("/users", CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{userId}", UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{userId}", DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/users/{userId}/worklogs", GetUserWorklogsHandler).Methods("GET")
	router.HandleFunc("/users/{userId}/tasks/{taskId}/start", StartTaskHandler).Methods("POST")
	router.HandleFunc("/users/{userId}/tasks/{taskId}/stop", StopTaskHandler).Methods("POST")
	
	// router.HandleFunc("/info", GetInfoHandler).Methods("GET")
}

func EncodePassportHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Passport encoded"))
}
