package persistence_impl

import (
	"context"

	"github.com/google/uuid"
	"github.com/little-tonii/gofiber-base/internal/domain/entity"
	"gorm.io/gorm"
)

type UserPersistenceImpl struct {
	db *gorm.DB
}

func NewUserPersistenceImpl(db *gorm.DB) *UserPersistenceImpl {
	return &UserPersistenceImpl{db: db}
}

func (persistence *UserPersistenceImpl) FindById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	result := new(entity.User)
	err := persistence.db.WithContext(ctx).
		Raw(`
			SELECT * FROM users WHERE id = @id
		`, map[string]any{
			"id": id,
		}).Scan(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (persistence *UserPersistenceImpl) Save(ctx context.Context, user *entity.User) error {
	err := persistence.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (persistence *UserPersistenceImpl) Update(ctx context.Context, user *entity.User) error {
	err := persistence.db.WithContext(ctx).Save(user).Error
	if err != nil {
		return err
	}
	return nil
}