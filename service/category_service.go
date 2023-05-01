package service

import (
	"context"

	"github.com/faridlan/nostra-api-product/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateReq) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateReq) web.CategoryResponse
	Delete(ctx context.Context, categoryId string)
	FindById(ctx context.Context, categoryId string) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
	CreateMany(ctx context.Context, request []web.CategoryCreateReq) []web.CategoryResponse
	DeleteAll(ctx context.Context)
}
