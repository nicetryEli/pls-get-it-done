package user_usecase

import (
	"context"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, req *RegisterUserReq) (*RegisterUserResp, error)
	LoginUser(ctx context.Context, req *LoginUserReq) (*LoginUserResp, error)
}
