package service

import (
	"context"

	"github.com/faridlan/nostra-api-product/model/web"
)

type RoleService interface {
	Create(ctx context.Context, request web.RoleCreateReq) web.RoleResponse
	Update(ctx context.Context, request web.RoleUpdateReq) web.RoleResponse
	FindById(ctx context.Context, roleId string) web.RoleResponse
	FindAll(ctx context.Context) []web.RoleResponse
	CreateMany(ctx context.Context, requests []web.RoleCreateReq) []web.RoleResponse
	DeleteAll(ctx context.Context)
}
