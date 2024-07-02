package repositories

import (
	"Effective_Mobile/internal/models"
	db "Effective_Mobile/internal/queries"
	"context"
	"database/sql"
	"log"
	"time"
)

type WorklogRepository struct {
	Queries *db.Queries
}

func NewWorklogRepository(dbConn *sql.DB) *WorklogRepository {
	return &WorklogRepository{
		Queries: db.New(dbConn),
	}
}

func (r *WorklogRepository) StartTask(worklog *models.Worklog) error {
	_, err := r.Queries.GetUserByID(context.Background(), worklog.UserID)
	if err != nil {
		log.Printf("failed to get user by ID: %v", err)
		return err
	}

	startedWorklog, err := r.Queries.StartTask(context.Background(), db.StartTaskParams{
		UserID:      worklog.UserID,
		Title:       worklog.Title,
		Description: sql.NullString{String: worklog.Description, Valid: worklog.Description != ""},
	})
	if err != nil {
		log.Printf("failed to start task: %v", err)
		return err
	}

	*worklog = models.Worklog{
		WorklogID:   startedWorklog.WorklogID,
		UserID:      startedWorklog.UserID,
		Title:       startedWorklog.Title,
		Description: startedWorklog.Description.String,
		StartTime:   startedWorklog.StartTime,
	}
	return nil
}

func (r *WorklogRepository) StopTask(worklog *models.Worklog) error {
	params := db.StopTaskParams{
		UserID:    worklog.UserID,
		WorklogID: worklog.WorklogID,
	}
	stoppedWorklog, err := r.Queries.StopTask(context.Background(), params)
	if err != nil {
		log.Printf("failed to stop task: %v", err)
		return err
	}

	*worklog = models.Worklog{
		WorklogID:   stoppedWorklog.WorklogID,
		UserID:      stoppedWorklog.UserID,
		Title:       stoppedWorklog.Title,
		Description: stoppedWorklog.Description.String,
		StartTime:   stoppedWorklog.StartTime,
		EndTime:     stoppedWorklog.EndTime.Time,
		HoursSpent:  stoppedWorklog.HoursSpent.Int64,
	}
	return nil
}

func (r *WorklogRepository) GetUserWorklogs(userID int32) ([]models.Worklog, error) {
	dbWorklogs, err := r.Queries.GetUserWorklogs(context.Background(), userID)
	if err != nil {
		log.Printf("failed to get user worklogs: %v", err)
		return nil, err
	}

	var worklogs []models.Worklog
	for _, dbWorklog := range dbWorklogs {
		endTime := time.Time{}
		if dbWorklog.EndTime.Valid {
			endTime = dbWorklog.EndTime.Time
		}

		hoursSpent := int64(0)
		if dbWorklog.HoursSpent.Valid {
			hoursSpent = dbWorklog.HoursSpent.Int64
		}

		worklogs = append(worklogs, models.Worklog{
			WorklogID:   dbWorklog.WorklogID,
			UserID:      dbWorklog.UserID,
			Title:       dbWorklog.Title,
			Description: dbWorklog.Description.String,
			StartTime:   dbWorklog.StartTime,
			EndTime:     endTime,
			HoursSpent:  hoursSpent,
		})
	}
	return worklogs, nil
}
