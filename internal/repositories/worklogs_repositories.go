package repositories

import (
	"context"
	"log"

	models "Effective_Mobile/internal/queries"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WorklogRepository struct {
	Queries *models.Queries
}

func NewWorklogRepository(queriesConn *pgxpool.Pool) (*WorklogRepository, error) {
	return &WorklogRepository{
		Queries: models.New(queriesConn),
	}, nil
}

func (r *WorklogRepository) StartTask(worklog *models.Worklog) error {
	user, err := r.Queries.GetUserByID(context.Background(), worklog.UserID)
	if err != nil {
		log.Printf("failed to get user by ID: %v", err)
		return err
	}
	log.Printf("Retrieved user: %+v", user)

	params := models.StartTaskParams{
		UserID:      worklog.UserID,
		Title:       worklog.Title,
		Description: worklog.Description,
	}
	startedWorklog, err := r.Queries.StartTask(context.Background(), params)
	if err != nil {
		log.Printf("failed to start task: %v", err)
		return err
	}
	*worklog = startedWorklog
	return nil
}

func (r *WorklogRepository) StopTask(worklog *models.Worklog) error {
	params := models.StopTaskParams{
		UserID:    worklog.UserID,
		WorklogID: worklog.WorklogID,
	}
	stoppedWorklog, err := r.Queries.StopTask(context.Background(), params)
	if err != nil {
		log.Printf("failed to stop task: %v", err)
		return err
	}
	*worklog = stoppedWorklog
	return nil
}

func (r *WorklogRepository) GetUserWorklogs(userID int32) ([]models.Worklog, error) {
	queriesWorklogs, err := r.Queries.GetUserWorklogs(context.Background(), userID)
	if err != nil {
		log.Printf("failed to get user worklogs: %v", err)
		return nil, err
	}
	return queriesWorklogs, nil
}
