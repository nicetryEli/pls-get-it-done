package persistence_impl

import (
	"context"

	"github.com/google/uuid"
	"github.com/little-tonii/gofiber-base/internal/domain/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserPersistenceImpl struct {
	db *gorm.DB
}

func NewUserPersistenceImpl(db *gorm.DB) *UserPersistenceImpl {
	return &UserPersistenceImpl{db: db}
}

func (persistence *UserPersistenceImpl) FindById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	result := new(entity.User)
	if err := persistence.db.WithContext(ctx).First(result, id).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (persistence *UserPersistenceImpl) Save(ctx context.Context, user *entity.User) error {
	if err := persistence.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (persistence *UserPersistenceImpl) Update(ctx context.Context, user *entity.User) error {
	if err := persistence.db.WithContext(ctx).
		Clauses(clause.Locking{Strength: "SHARE", Options: "NOWAIT"}).
		Save(user).Error; err != nil {
		return err
	}
	return nil
}
