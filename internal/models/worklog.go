package models

import "time"

type Worklog struct {
	WorklogID  int       `json:"worklog_id"`
	UserID     int       `json:"user_id"`
	TaskID     int       `json:"task_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	HoursSpent float64   `json:"hours_spent"`
}
