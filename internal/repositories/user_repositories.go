package repositories

import (
	"context"
	"database/sql"
	"log"

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

func (r *UserRepository) CreateUser(user *db.User) error {
	params := db.CreateUserParams{
		PassportSeries: user.PassportSeries,
		PassportNumber: user.PassportNumber,
		Surname:        user.Surname,
		Name:           user.Name,
		Patronymic:     user.Patronymic,
		Address:        user.Address,
	}
	createdUser, err := r.Queries.CreateUser(context.Background(), params)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return err
	}
	*user = createdUser
	return nil
}

func (r *UserRepository) GetUserByID(userID int32) (db.User, error) {
	user, err := r.Queries.GetUserByID(context.Background(), userID)
	if err != nil {
		log.Printf("failed to get user by ID: %v", err)
		return db.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetAllUsers() ([]db.User, error) {
	users, err := r.Queries.GetUsers(context.Background())
	if err != nil {
		log.Printf("failed to get all users: %v", err)
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(user *db.User) error {
	params := db.UpdateUserParams{
		UserID:         user.UserID,
		PassportSeries: user.PassportSeries,
		PassportNumber: user.PassportNumber,
		Surname:        user.Surname,
		Name:           user.Name,
		Patronymic:     user.Patronymic,
		Address:        user.Address,
	}
	updatedUser, err := r.Queries.UpdateUser(context.Background(), params)
	if err != nil {
		log.Printf("failed to update user: %v", err)
		return err
	}
	*user = updatedUser
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
