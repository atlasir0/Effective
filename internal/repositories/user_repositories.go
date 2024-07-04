package repositories

import (
	"context"
	"log"

	models "Effective_Mobile/internal/queries"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	Queries *models.Queries
}

func NewUserRepository(dbConn *pgxpool.Pool) (*UserRepository, error) {
	return &UserRepository{
		Queries: models.New(dbConn),
	}, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	params := models.CreateUserParams{
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
		Patronymic:     user.Patronymic,
		Address:        user.Address,
	}, nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	dbUsers, err := r.Queries.GetUsers(context.Background())
	if err != nil {
		log.Printf("failed to get all users: %v", err)
		return nil, err
	}
	return dbUsers, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	params := models.UpdateUserParams{
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

func (r *UserRepository) GetPaginatedUsers(limit, offset int32) ([]models.User, error) {
	params := models.GetPaginatedUsersParams{
		Limit:  limit,
		Offset: offset,
	}
	dbUsers, err := r.Queries.GetPaginatedUsers(context.Background(), params)
	if err != nil {
		log.Printf("failed to get paginated users: %v", err)
		return nil, err
	}
	return dbUsers, nil
}

func (r *UserRepository) GetFilteredUsers(column1, column2 string) ([]models.User, error) {
	log.Printf("Filtering users by column1: %s, column2: %s", column1, column2)
	params := models.GetFilteredUsersParams{
		Column1: column1,
		Surname: column2,
	}
	dbUsers, err := r.Queries.GetFilteredUsers(context.Background(), params)
	if err != nil {
		log.Printf("failed to get filtered users: %v", err)
		return nil, err
	}
	log.Printf("Found %d users", len(dbUsers))
	return dbUsers, nil
}
