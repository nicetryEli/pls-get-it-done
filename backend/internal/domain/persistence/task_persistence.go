package persistence

import (
	"context"

	"github.com/google/uuid"
	"github.com/little-tonii/gofiber-base/internal/domain/entity"
)

type TaskPersistence interface {
	FindById(ctx context.Context, id uuid.UUID) (*entity.Task, error)
	Save(ctx context.Context, task *entity.Task) error
	Update(ctx context.Context, task *entity.Task) error
}
