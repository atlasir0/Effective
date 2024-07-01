package services

import (
	"Effective_Mobile/internal/repositories"
	"Effective_Mobile/internal/models"
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

func (s *WorklogService) StartTask(worklog *models.Worklog) error {
	return s.WorklogRepo.StartTask(worklog)
}

func (s *WorklogService) StopTask(worklog *models.Worklog) error {
	return s.WorklogRepo.StopTask(worklog)
}

func (s *WorklogService) GetUserWorklogs(userID int32) ([]models.Worklog, error) {
	return s.WorklogRepo.GetUserWorklogs(userID)
}
