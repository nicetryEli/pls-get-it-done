package provider

import (
	"context"

	"gorm.io/gorm"
)

type TransactionProvider interface {
	Create(ctx context.Context) *gorm.DB
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
	Transaction(ctx context.Context, txFunc func(tx *gorm.DB) error) error
}
