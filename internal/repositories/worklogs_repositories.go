package repositories

import (
	"context"
	"database/sql"
	"log"

	db "Effective_Mobile/internal/queries"
)

type WorklogRepository struct {
	Queries *db.Queries
}

func NewWorklogRepository(dbConn *sql.DB) *WorklogRepository {
	return &WorklogRepository{
		Queries: db.New(dbConn),
	}
}

func (r *WorklogRepository) StartTask(worklog *db.Worklog) error {
	params := db.StartTaskParams{
		UserID: worklog.UserID,
		TaskID: worklog.TaskID,
	}
	startedWorklog, err := r.Queries.StartTask(context.Background(), params)
	if err != nil {
		log.Printf("failed to start task: %v", err)
		return err
	}
	*worklog = startedWorklog
	return nil
}

func (r *WorklogRepository) StopTask(worklog *db.Worklog) error {
	params := db.StopTaskParams{
		UserID: worklog.UserID,
		TaskID: worklog.TaskID,
	}
	stoppedWorklog, err := r.Queries.StopTask(context.Background(), params)
	if err != nil {
		log.Printf("failed to stop task: %v", err)
		return err
	}
	*worklog = stoppedWorklog
	return nil
}

func (r *WorklogRepository) GetUserWorklogs(userID int32) ([]db.Worklog, error) {
	worklogs, err := r.Queries.GetUserWorklogs(context.Background(), userID)
	if err != nil {
		log.Printf("failed to get user worklogs: %v", err)
		return nil, err
	}
	return worklogs, nil
}