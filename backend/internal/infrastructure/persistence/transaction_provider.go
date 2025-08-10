package persistence_impl

import (
	"context"

	"gorm.io/gorm"
)

type TransactionProviderImpl struct {
	db *gorm.DB
}

func NewTransactionProviderImpl(db *gorm.DB) *TransactionProviderImpl {
	return &TransactionProviderImpl{
		db: db,
	}
}

func (persistence *TransactionProviderImpl) Create(ctx context.Context) *gorm.DB {
	return persistence.db.WithContext(ctx).Begin()
}

func (persistence *TransactionProviderImpl) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (persistence *TransactionProviderImpl) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (persistence *TransactionProviderImpl) Transaction(ctx context.Context, txFunc func(tx *gorm.DB) error) error {
	return persistence.db.WithContext(ctx).Transaction(txFunc)
}
