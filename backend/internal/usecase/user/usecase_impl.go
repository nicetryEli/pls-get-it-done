package user_usecase

import (
	"context"

	"github.com/little-tonii/gofiber-base/internal/domain/persistence"
)

type UserUsecaseImpl struct {
	userPersis persistence.UserPersistence
	txProvider persistence.TransactionProvider
}

func NewUserUsecaseImpl(userPersis persistence.UserPersistence, txProvider persistence.TransactionProvider) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		userPersis: userPersis,
		txProvider: txProvider,
	}
}

func (usecase *UserUsecaseImpl) RegisterUser(ctx context.Context, req *RegisterUserReq) (*RegisterUserResp, error) {
	return nil, nil
}

func (usecase *UserUsecaseImpl) LoginUser(ctx context.Context, req *LoginUserReq) (*LoginUserResp, error) {
	return nil, nil
}
