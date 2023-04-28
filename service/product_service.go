package service

import (
	"context"

	"github.com/faridlan/nostra-api-product/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateReq) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateReq) web.ProductResponse
	Delete(ctx context.Context, productId string)
	FindById(ctx context.Context, productId string) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
	CreateMany(ctx context.Context, request []web.ProductCreateReq) []web.ProductResponse
	DeleteAll(ctx context.Context)
}
