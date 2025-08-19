package persistence

import (
	"context"

	"github.com/google/uuid"
	"github.com/little-tonii/gofiber-base/internal/domain/entity"
)

type PomodoroPersistence interface {
	FindById(ctx context.Context, id uuid.UUID) (*entity.Pomodoro, error)
	Save(ctx context.Context, pomodoro *entity.Pomodoro) error
	Update(ctx context.Context, pomodoro *entity.Pomodoro) error
}
