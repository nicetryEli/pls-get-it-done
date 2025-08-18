package persistence_impl

import (
	"context"

	"github.com/google/uuid"
	"github.com/little-tonii/gofiber-base/internal/domain/entity"
	"gorm.io/gorm"
)

type PomodoroPersistenceImpl struct {
	db *gorm.DB
}

func NewPomodoroPersistenceImpl(db *gorm.DB) *PomodoroPersistenceImpl {
	return &PomodoroPersistenceImpl{db: db}
}

func (persistence *PomodoroPersistenceImpl) FindById(ctx context.Context, id uuid.UUID) (*entity.Pomodoro, error) {
	result := new(entity.Pomodoro)
	err := persistence.db.WithContext(ctx).
		Raw(`
			SELECT * FROM pomodoros WHERE id = @id
		`, map[string]any{
			"id": id,
		}).Scan(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (persistence *PomodoroPersistenceImpl) Save(ctx context.Context, pomodoro *entity.Pomodoro) error {
	err := persistence.db.WithContext(ctx).Create(pomodoro).Error
	if err != nil {
		return err
	}
	return nil
}

func (persistence *PomodoroPersistenceImpl) Update(ctx context.Context, pomodoro *entity.Pomodoro) error {
	err := persistence.db.WithContext(ctx).Save(pomodoro).Error
	if err != nil {
		return err
	}
	return nil
}
