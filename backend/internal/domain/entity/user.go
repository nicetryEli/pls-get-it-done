package entity

import (
	"time"

	"github.com/google/uuid"
)

var (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

type User struct {
	Id                uuid.UUID `gorm:"column:id;primaryKey"`
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt         time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
	Username          string    `gorm:"column:username"`
	FullName          string    `gorm:"column:full_name"`
	Email             string    `gorm:"column:email"`
	TokenVersion      int64     `gorm:"column:token_version"`
	IsActive          bool      `gorm:"column:is_active"`
	IsVerified        bool      `gorm:"column:is_verified"`
	HashedPassword    string    `gorm:"column:hashed_password"`
	LastActivityAt    time.Time `gorm:"column:last_activity_at;autoCreateTime:milli"`
	PasswordUpdatedAt time.Time `gorm:"column:password_updated_at;autoCreateTime:milli"`
	Role              string    `gorm:"column:role"`
	AvatarKey         *string   `gorm:"column:avatar_key"`
}

func (User) TableName() string {
	return "users"
}
