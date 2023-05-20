package service

import (
	"context"

	"github.com/faridlan/nostra-api-product/model/web"
)

type AuthService interface {
	Register(ctx context.Context, request web.UserCreateReq) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateReq) web.UserResponse
	FindById(ctx context.Context, userId string) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
	Login(ctx context.Context, request web.Login) web.LoginResponse
	SaveMany(ctx context.Context, request []web.UserCreateReq) []web.UserResponse
	DeleteAll(ctx context.Context)
}
