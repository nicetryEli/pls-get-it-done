package persistence_impl

import (
	"context"

	"github.com/google/uuid"
	"github.com/little-tonii/gofiber-base/internal/domain/entity"
	"gorm.io/gorm"
)

type TodoPersistenceImpl struct {
	db *gorm.DB
}

func NewTodoPersistenceImpl(db *gorm.DB) *TodoPersistenceImpl {
	return &TodoPersistenceImpl{db: db}
}

func (persistence *TodoPersistenceImpl) FindById(ctx context.Context, id uuid.UUID) (*entity.Todo, error) {
	result := new(entity.Todo)
	err := persistence.db.WithContext(ctx).
		Raw(`
			SELECT * FROM todos WHERE id = @id
		`, map[string]any{
			"id": id,
		}).Scan(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (persistence *TodoPersistenceImpl) Save(ctx context.Context, todo *entity.Todo) error {
	err := persistence.db.WithContext(ctx).Create(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (persistence *TodoPersistenceImpl) Update(ctx context.Context, todo *entity.Todo) error {
	err := persistence.db.WithContext(ctx).Save(todo).Error
	if err != nil {
		return err
	}
	return nil
}
