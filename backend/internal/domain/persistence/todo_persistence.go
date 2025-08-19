package persistence

import (
	"context"

	"github.com/google/uuid"
	"github.com/little-tonii/gofiber-base/internal/domain/entity"
)

type TodoPersistence interface {
	FindById(ctx context.Context, id uuid.UUID) (*entity.Todo, error)
	Save(ctx context.Context, todo *entity.Todo) error
	Update(ctx context.Context, todo *entity.Todo) error
}
