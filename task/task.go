package task

import "time"

type (
	Task struct {
		Id          int
		Description string
		Status      string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)

func NewTask(description string, date time.Time) *Task {
	return &Task{
		Id:          1,
		Description: description,
		Status:      "todo",
		CreatedAt:   date,
		UpdatedAt:   date,
	}
}
