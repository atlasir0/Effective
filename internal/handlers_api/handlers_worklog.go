package handlers

import (
	db "Effective_Mobile/internal/queries" //Это модель а не бд!
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
)

// @Summary Start a task
// @Description Start a new task
// @Tags worklogs
// @Accept  json
// @Produce  json
// @Param worklog body db.Worklog true "Worklog object"
// @Success 201 {object} db.Worklog
// @Router /worklogs [post]
func (h *WorklogHandler) StartTask(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to start task")

	var worklog db.Worklog
	err := json.NewDecoder(r.Body).Decode(&worklog)
	if err != nil {
		log.Printf("Error decoding start task request body: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if worklog.UserID == 0 {
		log.Println("UserID is 0, invalid request")
		http.Error(w, "Invalid UserID", http.StatusBadRequest)
		return
	}

	startTime := time.Now()
	worklog.StartTime = pgtype.Timestamp{Time: startTime, Valid: true}

	if err := h.WorklogService.StartTask(&worklog); err != nil {
		log.Printf("Error starting task: %v", err)
		http.Error(w, "Failed to start task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "200"}
	json.NewEncoder(w).Encode(response)
	log.Println("Successfully started task and sent response")
}

// @Summary Stop a task
// @Description Stop a task by ID
// @Tags worklogs
// @Accept  json
// @Produce  json
// @Param id path int true "Worklog ID"
// @Param worklog body db.Worklog true "Worklog object"
// @Success 200 {object} db.Worklog
// @Router /worklogs/{id} [put]
func (h *WorklogHandler) StopTask(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to stop task")

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var worklog db.Worklog
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "200"}
	json.NewEncoder(w).Encode(response)
	log.Println("Successfully stopped task and sent 200")
}

// @Summary Get user worklogs
// @Description Get a list of worklogs for a user by user ID
// @Tags worklogs
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {array} db.Worklog
// @Router /worklogs/user/{id} [get]
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
		worklogs[i].StartTime.Time = worklogs[i].StartTime.Time.Truncate(time.Minute)
		worklogs[i].EndTime.Time = worklogs[i].EndTime.Time.Truncate(time.Minute)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(worklogs)
	log.Println("Successfully retrieved user worklogs and sent response")
}
