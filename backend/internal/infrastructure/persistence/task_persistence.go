package persistence_impl

import (
	"context"

	"github.com/google/uuid"
	"github.com/little-tonii/gofiber-base/internal/domain/entity"
	"gorm.io/gorm"
)

type TaskPersistenceImpl struct {
	db *gorm.DB
}

func NewTaskPersistenceImpl(db *gorm.DB) *TaskPersistenceImpl {
	return &TaskPersistenceImpl{db: db}
}

func (persistence *TaskPersistenceImpl) FindById(ctx context.Context, id uuid.UUID) (*entity.Task, error) {
	result := new(entity.Task)
	err := persistence.db.WithContext(ctx).
		Raw(`
			SELECT * FROM tasks WHERE id = @id
		`, map[string]any{
			"id": id,
		}).Scan(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// func (persistence *TaskPersistenceImpl) FindByName(ctx context.Context, name string) (*entity.Task, error) {
// 	result := new(entity.Task)
// 	err := persistence.db.WithContext(ctx).
// 		Raw(`
// 			SELECT * FROM tasks WHERE name = @name
// 		`, map[string]any{
// 			"name": name,
// 		}).Scan(result).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

func (persistence *TaskPersistenceImpl) Save(ctx context.Context, task *entity.Task) error {
	err := persistence.db.WithContext(ctx).Create(task).Error
	if err != nil {
		return err
	}
	return nil
}

func (persistence *TaskPersistenceImpl) Update(ctx context.Context, task *entity.Task) error {
	err := persistence.db.WithContext(ctx).Save(task).Error
	if err != nil {
		return err
	}
	return nil
}
