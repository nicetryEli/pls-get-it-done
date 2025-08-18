package provider_impl

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

func (provider *TransactionProviderImpl) Create(ctx context.Context) *gorm.DB {
	return provider.db.WithContext(ctx).Begin()
}

func (provider *TransactionProviderImpl) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (provider *TransactionProviderImpl) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (provider *TransactionProviderImpl) Transaction(ctx context.Context, txFunc func(tx *gorm.DB) error) error {
	return provider.db.WithContext(ctx).Transaction(txFunc)
}
