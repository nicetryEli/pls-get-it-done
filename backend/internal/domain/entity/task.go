package entity

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

var (
	Processing TaskStatus = "processing"
	Completed  TaskStatus = "completed"
	Delayed    TaskStatus = "delayed"
	Cancelled  TaskStatus = "cancelled"
)

type Task struct {
	Id         uuid.UUID  `gorm:"column:id;primaryKey"`
	TodoId     uuid.UUID  `gorm:"column:todo_id"`
	PomodoroId uuid.UUID  `gorm:"column:pomodoro_id"`
	Name       string     `gorm:"column:name"`
	Status     TaskStatus `gorm:"column:status"`
	CreatedAt  time.Time  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt  time.Time  `gorm:"column:updated_at;autoUpdateTime:milli"`
}

func (Task) TableName() string {
	return "tasks"
}
