package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPaginatedUsersHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	page, _ := strconv.Atoi(queryParams.Get("page"))
	pageSize, _ := strconv.Atoi(queryParams.Get("pageSize"))

	users, err := getPaginatedUsersFromDB(page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getPaginatedUsersFromDB(page, pageSize int) ([]map[string]interface{}, error) {
	return []map[string]interface{}{
		{"id": 1, "name": "John Doe"},
		{"id": 2, "name": "Jane Doe"},
	}, nil
}

func GetFilteredUsersHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	filterField := queryParams.Get("filterField")
	filterValue := queryParams.Get("filterValue")

	users, err := getFilteredUsersFromDB(filterField, filterValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getFilteredUsersFromDB(filterField, filterValue string) ([]map[string]interface{}, error) {

	return []map[string]interface{}{
		{"id": 1, "name": "John Doe"},
		{"id": 2, "name": "Jane Doe"},
	}, nil
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	user, err := getUserByIDFromDB(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getUserByIDFromDB(userId string) (map[string]interface{}, error) {

	return map[string]interface{}{
		"id": 1, "name": "John Doe",
	}, nil
}

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
		Surname        string `json:"surname"`
		Name           string `json:"name"`
		Patronymic     string `json:"patronymic"`
		Address        string `json:"address"`
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	newUser, err := createUserInDB(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func createUserInDB(user struct {
	PassportNumber string `json:"passportNumber"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"user_id":        1,
		"passportNumber": user.PassportNumber,
		"surname":        user.Surname,
		"name":           user.Name,
		"patronymic":     user.Patronymic,
		"address":        user.Address,
	}, nil
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
