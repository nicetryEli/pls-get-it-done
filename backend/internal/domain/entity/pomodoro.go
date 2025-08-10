package entity

import (
	"time"

	"github.com/google/uuid"
)

type Pomodoro struct {
	Id            uuid.UUID `gorm:"column:id;primaryKey"`
	Repeat        int       `gorm:"column:repeat"`
	Duration      int       `gorm:"column:duration"`
	BreakDuration int       `gorm:"column:break_duration"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
}

func (Pomodoro) TableName() string {
	return "pomodoros"
}
