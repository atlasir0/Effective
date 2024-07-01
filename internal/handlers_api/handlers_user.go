package handlers

import (
	"Effective_Mobile/internal/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
//TODO: выводить статус 200 в случае успешного результата
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to create user")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if len(bodyBytes) == 0 {
		log.Println("Request body is empty")
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.Unmarshal(bodyBytes, &user); err != nil {
		log.Printf("Invalid JSON: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.UserService.CreateUser(&user); err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	log.Println("Successfully created user")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200"))
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get all users")

	users, err := h.UserService.GetAllUsers()
	if err != nil {
		log.Printf("Error getting users: %v", err)
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get user by ID")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.UserService.GetUserByID(int32(id))
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to update user")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if len(bodyBytes) == 0 {
		log.Println("Request body is empty")
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.Unmarshal(bodyBytes, &user); err != nil {
		log.Printf("Invalid JSON: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user.UserID = int32(id)

	if err := h.UserService.UpdateUser(&user); err != nil {
		log.Printf("Error updating user: %v", err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("Successfully updated user")
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to delete user")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.UserService.DeleteUser(int32(id)); err != nil {
		log.Printf("Error deleting user: %v", err)
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("Successfully deleted user")
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
