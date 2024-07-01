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
	params := db.StartTaskParams{
		UserID: worklog.UserID,
		TaskID: worklog.TaskID,
	}
	startedWorklog, err := r.Queries.StartTask(context.Background(), params)
	if err != nil {
		log.Printf("failed to start task: %v", err)
		return err
	}

	endTime := time.Time{}
	if startedWorklog.EndTime.Valid {
		endTime = startedWorklog.EndTime.Time
	}

	hoursSpent := int64(0)
	if startedWorklog.HoursSpent.Valid {
		hoursSpent = startedWorklog.HoursSpent.Int64
	}

	*worklog = models.Worklog{
		WorklogID:  startedWorklog.WorklogID,
		UserID:     startedWorklog.UserID,
		TaskID:     startedWorklog.TaskID,
		StartTime:  startedWorklog.StartTime,
		EndTime:    endTime,
		HoursSpent: hoursSpent,
	}
	return nil
}

func (r *WorklogRepository) StopTask(worklog *models.Worklog) error {
	params := db.StopTaskParams{
		UserID: worklog.UserID,
		TaskID: worklog.TaskID,
	}
	stoppedWorklog, err := r.Queries.StopTask(context.Background(), params)
	if err != nil {
		log.Printf("failed to stop task: %v", err)
		return err
	}

	endTime := time.Time{}
	if stoppedWorklog.EndTime.Valid {
		endTime = stoppedWorklog.EndTime.Time
	}

	hoursSpent := int64(0)
	if stoppedWorklog.HoursSpent.Valid {
		hoursSpent = stoppedWorklog.HoursSpent.Int64
	}

	*worklog = models.Worklog{
		WorklogID:  stoppedWorklog.WorklogID,
		UserID:     stoppedWorklog.UserID,
		TaskID:     stoppedWorklog.TaskID,
		StartTime:  stoppedWorklog.StartTime,
		EndTime:    endTime,
		HoursSpent: hoursSpent,
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
			WorklogID:  dbWorklog.WorklogID,
			UserID:     dbWorklog.UserID,
			TaskID:     dbWorklog.TaskID,
			StartTime:  dbWorklog.StartTime,
			EndTime:    endTime,
			HoursSpent: hoursSpent,
		})
	}
	return worklogs, nil
}
