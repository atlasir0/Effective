package api

import (
	db "Effective_Mobile/internal/queries"
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var q *db.Queries

func InitDB(database *sql.DB) {
	q = db.New(database)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to create user")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		log.Println("Request body is nil")
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	log.Printf("Request body: %s", string(bodyBytes))

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	var user db.CreateUserParams
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding create user request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser, err := q.CreateUser(r.Context(), user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
	log.Println("Successfully created user and sent response")
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get all users")
	if q == nil {
		log.Println("Queries instance is nil")
		http.Error(w, "Database not initialized", http.StatusInternalServerError)
		return
	}

	users, err := q.GetUsers(r.Context())
	if err != nil {
		log.Printf("Error getting all users: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
	log.Println("Successfully sent all users response")
}

/////////////////////////////////////

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to update user")
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["userId"])

	var user db.UpdateUserParams
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding update user request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.UserID = int32(userId)
	updatedUser, err := q.UpdateUser(r.Context(), user)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
	log.Println("Successfully updated user and sent response")
}

func GetPaginatedUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get paginated users")
	queryParams := r.URL.Query()
	page, _ := strconv.Atoi(queryParams.Get("page"))
	pageSize, _ := strconv.Atoi(queryParams.Get("pageSize"))

	users, err := q.GetPaginatedUsers(r.Context(), db.GetPaginatedUsersParams{
		Limit:  int32(pageSize),
		Offset: int32((page - 1) * pageSize),
	})
	if err != nil {
		log.Printf("Error getting paginated users: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
	log.Println("Successfully sent paginated users response")
}

func GetFilteredUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get filtered users")
	queryParams := r.URL.Query()
	filterField := queryParams.Get("filterField")
	filterValue := queryParams.Get("filterValue")

	users, err := q.GetFilteredUsers(r.Context(), db.GetFilteredUsersParams{
		Column1: filterField,
		Column2: filterValue,
	})
	if err != nil {
		log.Printf("Error getting filtered users: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
	log.Println("Successfully sent filtered users response")
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get user by ID")
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["userId"])

	user, err := q.GetUserByID(r.Context(), int32(userId))
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	log.Println("Successfully sent user by ID response")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to delete user")
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["userId"])

	err := q.DeleteUser(r.Context(), int32(userId))
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Println("Successfully deleted user")
}

func GetUserWorklogsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get user worklogs")
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["userId"])

	worklogs, err := q.GetUserWorklogs(r.Context(), int32(userId))
	if err != nil {
		log.Printf("Error getting user worklogs: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(worklogs)
	log.Println("Successfully sent user worklogs response")
}

func StartTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to start task")
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["userId"])
	taskId, _ := strconv.Atoi(vars["taskId"])

	worklog, err := q.StartTask(r.Context(), db.StartTaskParams{
		UserID: int32(userId),
		TaskID: int32(taskId),
	})
	if err != nil {
		log.Printf("Error starting task: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(worklog)
	log.Println("Successfully started task and sent response")
}

func StopTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to stop task")
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["userId"])
	taskId, _ := strconv.Atoi(vars["taskId"])

	worklog, err := q.StopTask(r.Context(), db.StopTaskParams{
		UserID: int32(userId),
		TaskID: int32(taskId),
	})
	if err != nil {
		log.Printf("Error stopping task: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(worklog)
	log.Println("Successfully stopped task and sent response")
}
