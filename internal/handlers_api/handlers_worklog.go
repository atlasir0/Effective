package handlers

import (
	"Effective_Mobile/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (h *WorklogHandler) StartTask(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to start task")

	var worklog models.Worklog
	err := json.NewDecoder(r.Body).Decode(&worklog)
	if err != nil {
		log.Printf("Error decoding start task request body: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	worklog.StartTime = time.Now()

	if err := h.WorklogService.StartTask(&worklog); err != nil {
		log.Printf("Error starting task: %v", err)
		http.Error(w, "Failed to start task", http.StatusInternalServerError)
		return
	}

	worklog.StartTime = worklog.StartTime.Truncate(time.Minute)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(worklog)
	log.Println("Successfully started task and sent response")
}

func (h *WorklogHandler) StopTask(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to stop task")

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var worklog models.Worklog
	err = json.NewDecoder(r.Body).Decode(&worklog)
	if err != nil {
		log.Printf("Error decoding stop task request body: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	worklog.UserID = int32(userID)

	if err := h.WorklogService.StopTask(&worklog); err != nil {
		log.Printf("Error stopping task: %v", err)
		http.Error(w, "Failed to stop task", http.StatusInternalServerError)
		return
	}

	worklog.EndTime = worklog.EndTime.Truncate(time.Minute)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(worklog)
	log.Println("Successfully stopped task and sent response")
}

func (h *WorklogHandler) GetUserWorklogs(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get user worklogs")

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	worklogs, err := h.WorklogService.GetUserWorklogs(int32(userID))
	if err != nil {
		log.Printf("Error getting user worklogs: %v", err)
		http.Error(w, "Failed to get user worklogs", http.StatusInternalServerError)
		return
	}

	for i := range worklogs {
		worklogs[i].StartTime = worklogs[i].StartTime.Truncate(time.Minute)
		worklogs[i].EndTime = worklogs[i].EndTime.Truncate(time.Minute)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(worklogs)
	log.Println("Successfully retrieved user worklogs and sent response")
}
