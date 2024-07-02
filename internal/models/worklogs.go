package models

import (
	"time"
)

type Worklog struct {
	WorklogID   int32     `json:"worklog_id,omitempty"`
	UserID      int32     `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	StartTime   time.Time `json:"start_time,omitempty"`
	EndTime     time.Time `json:"end_time,omitempty"`
	HoursSpent  int64     `json:"hours_spent,omitempty"`
}

