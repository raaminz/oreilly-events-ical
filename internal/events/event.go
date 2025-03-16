package events

import "time"

type OreillyEvent struct {
	ID          string
	Topics      string
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Instructors string
	Levels      string
}
