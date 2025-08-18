package entity

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id        uuid.UUID `gorm:"column:id;primaryKey"`
	UserId    uuid.UUID `gorm:"column:user_id"`
	Title     string    `gorm:"column:title"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
}

func (Todo) TableName() string {
	return "todos"
}
