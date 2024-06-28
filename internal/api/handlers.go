package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	
	queryParams := r.URL.Query()
	filterField := queryParams.Get("filterField")
	page, _ := strconv.Atoi(queryParams.Get("page"))
	pageSize, _ := strconv.Atoi(queryParams.Get("pageSize"))

	_ = filterField
	_ = page
	_ = pageSize

	users := []map[string]interface{}{
		{"id": 1, "name": "John Doe"},
		{"id": 2, "name": "Jane Doe"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		PassportNumber string `json:"passportNumber"`
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser := map[string]interface{}{
		"id":             1,
		"passportNumber": user.PassportNumber,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	var user map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	user["id"] = userId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	response := map[string]interface{}{
		"id":      userId,
		"deleted": true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetUserWorklogsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	queryParams := r.URL.Query()
	startDate := queryParams.Get("startDate")
	endDate := queryParams.Get("endDate")
	sortBy := queryParams.Get("sortBy")
	order := queryParams.Get("order")


	_ = userId
	_ = startDate
	_ = endDate
	_ = sortBy
	_ = order

	worklogs := []map[string]interface{}{
		{"taskId": 1, "hoursSpent": 5},
		{"taskId": 2, "hoursSpent": 3},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(worklogs)
}

func StartTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	taskId := vars["taskId"]

	response := map[string]interface{}{
		"userId": userId,
		"taskId": taskId,
		"status": "started",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func StopTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	taskId := vars["taskId"]


	response := map[string]interface{}{
		"userId": userId,
		"taskId": taskId,
		"status": "stopped",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetInfoHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	passportSerie := queryParams.Get("passportSerie")
	passportNumber := queryParams.Get("passportNumber")


	info := map[string]interface{}{
		"passportSerie":  passportSerie,
		"passportNumber": passportNumber,
		"additionalInfo": "Some additional info",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}
