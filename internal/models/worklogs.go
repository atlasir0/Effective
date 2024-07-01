package models

import (
	"time"
)

type Worklog struct {
	WorklogID  int32     `json:"worklog_id"`
	UserID     int32     `json:"user_id"`
	TaskID     int32     `json:"task_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	HoursSpent int64     `json:"hours_spent"`
}

type Task struct {
	TaskID      int32  `json:"task_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
