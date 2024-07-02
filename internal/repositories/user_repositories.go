package repositories

import (
	"context"
	"database/sql"
	"log"

	"Effective_Mobile/internal/models"    
	db "Effective_Mobile/internal/queries" 
)

type UserRepository struct {
	Queries *db.Queries
}

func NewUserRepository(dbConn *sql.DB) *UserRepository {
	return &UserRepository{
		Queries: db.New(dbConn),
	}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	params := db.CreateUserParams{
		PassportSeries: user.PassportSeries,
		PassportNumber: user.PassportNumber,
		Surname:        user.Surname,
		Name:           user.Name,
		Patronymic:     sql.NullString{String: user.Patronymic, Valid: user.Patronymic != ""},
		Address:        sql.NullString{String: user.Address, Valid: user.Address != ""},
	}
	createdUser, err := r.Queries.CreateUser(context.Background(), params)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return err
	}
	*user = models.User{
		UserID:         createdUser.UserID,
		PassportSeries: createdUser.PassportSeries,
		PassportNumber: createdUser.PassportNumber,
		Surname:        createdUser.Surname,
		Name:           createdUser.Name,
		Patronymic:     createdUser.Patronymic.String,
		Address:        createdUser.Address.String,
	}
	log.Printf("user created successfully: %+v", createdUser)
	return nil
}

func (r *UserRepository) GetUserByID(userID int32) (models.User, error) {
	user, err := r.Queries.GetUserByID(context.Background(), userID)
	if err != nil {
		log.Printf("failed to get user by ID: %v", err)
		return models.User{}, err
	}
	return models.User{
		UserID:         user.UserID,
		PassportSeries: user.PassportSeries,
		PassportNumber: user.PassportNumber,
		Surname:        user.Surname,
		Name:           user.Name,
		Patronymic:     user.Patronymic.String,
		Address:        user.Address.String,
	}, nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	dbUsers, err := r.Queries.GetUsers(context.Background())
	if err != nil {
		log.Printf("failed to get all users: %v", err)
		return nil, err
	}
	var users []models.User
	for _, dbUser := range dbUsers {
		users = append(users, models.User{
			UserID:         dbUser.UserID,
			PassportSeries: dbUser.PassportSeries,
			PassportNumber: dbUser.PassportNumber,
			Surname:        dbUser.Surname,
			Name:           dbUser.Name,
			Patronymic:     dbUser.Patronymic.String,
			Address:        dbUser.Address.String,
		})
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	params := db.UpdateUserParams{
		UserID:         user.UserID,
		PassportSeries: user.PassportSeries,
		PassportNumber: user.PassportNumber,
		Surname:        user.Surname,
		Name:           user.Name,
		Patronymic:     sql.NullString{String: user.Patronymic, Valid: user.Patronymic != ""},
		Address:        sql.NullString{String: user.Address, Valid: user.Address != ""},
	}
	updatedUser, err := r.Queries.UpdateUser(context.Background(), params)
	if err != nil {
		log.Printf("failed to update user: %v", err)
		return err
	}
	*user = models.User{
		UserID:         updatedUser.UserID,
		PassportSeries: updatedUser.PassportSeries,
		PassportNumber: updatedUser.PassportNumber,
		Surname:        updatedUser.Surname,
		Name:           updatedUser.Name,
		Patronymic:     updatedUser.Patronymic.String,
		Address:        updatedUser.Address.String,
	}
	return nil
}

func (r *UserRepository) DeleteUser(userID int32) error {
	err := r.Queries.DeleteUser(context.Background(), userID)
	if err != nil {
		log.Printf("failed to delete user: %v", err)
		return err
	}
	return nil 
}
