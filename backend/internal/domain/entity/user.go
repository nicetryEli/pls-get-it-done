package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID `gorm:"column:id;primaryKey"`
	Name         string    `gorm:"column:username"`
	Avatar       string    `gorm:"column:avatar"`
	HashPassword string    `gorm:"column:hash_password"`
	Email        string    `gorm:"column:email;uniqueIndex"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
}

func (User) TableName() string {
	return "users"
}
