package handlers

import (
	models "Effective_Mobile/internal/queries"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "200", "message": "User created"})
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

	log.Println("Successfully updated user")
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "200", "message": "User updated successfully"})
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

	log.Println("Successfully deleted user")
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "200", "message": "User deleted successfully"})
}

func (h *UserHandler) GetPaginatedUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get paginated users")

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	users, err := h.UserService.GetPaginatedUsers(int32(limit), int32(offset))
	if err != nil {
		log.Printf("Error getting paginated users: %v", err)
		http.Error(w, "Failed to get paginated users", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetFilteredUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get filtered users")

	column1 := r.URL.Query().Get("column1")
	column2 := r.URL.Query().Get("column2")

	if column1 == "" || column2 == "" {
		http.Error(w, "Both column1 and column2 parameters are required", http.StatusBadRequest)
		return
	}

	users, err := h.UserService.GetFilteredUsers(column1, column2)
	if err != nil {
		log.Printf("Error getting filtered users: %v", err)
		http.Error(w, "Failed to get filtered users", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}