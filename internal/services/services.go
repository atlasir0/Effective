package services

import (
	models "Effective_Mobile/internal/queries"
	"Effective_Mobile/internal/repositories"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

type WorklogService struct {
	WorklogRepo *repositories.WorklogRepository
}

func NewWorklogService(worklogRepo *repositories.WorklogRepository) *WorklogService {
	return &WorklogService{
		WorklogRepo: worklogRepo,
	}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepo.CreateUser(user)
}

func (s *UserService) GetUserByID(userID int32) (models.User, error) {
	return s.UserRepo.GetUserByID(userID)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepo.GetAllUsers()
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.UserRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(userID int32) error {
	return s.UserRepo.DeleteUser(userID)
}

func (s *UserService) GetPaginatedUsers(limit, offset int32) ([]models.User, error) {
	return s.UserRepo.GetPaginatedUsers(limit, offset)
}

func (s *UserService) GetFilteredUsers(column1, column2 string) ([]models.User, error) {
	return s.UserRepo.GetFilteredUsers(column1, column2)
}

func (s *WorklogService) StartTask(worklog *models.Worklog) error {
	return s.WorklogRepo.StartTask(worklog)
}

func (s *WorklogService) StopTask(worklog *models.Worklog) error {
	return s.WorklogRepo.StopTask(worklog)
}

func (s *WorklogService) GetUserWorklogs(userID int32) ([]models.Worklog, error) {
	return s.WorklogRepo.GetUserWorklogs(userID)
}
