package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/domain"
	"github.com/faridlan/nostra-api-product/model/web"
	"github.com/faridlan/nostra-api-product/repository"
)

type CategoryServiceImpl struct {
	CategoryRepo repository.CategoryRepository
	DB           *sql.DB
}

func NewCategoryService(categoryRepo repository.CategoryRepository, db *sql.DB) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepo: categoryRepo,
		DB:           db,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateReq) web.CategoryResponse {
	panic("not implemented") // TODO: Implement
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateReq) web.CategoryResponse {
	panic("not implemented") // TODO: Implement
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId string) {
	panic("not implemented") // TODO: Implement
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId string) web.CategoryResponse {
	panic("not implemented") // TODO: Implement
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	panic("not implemented") // TODO: Implement
}

func (service *CategoryServiceImpl) CreateMany(ctx context.Context, request []web.CategoryCreateReq) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := []domain.Category{}

	for _, req := range request {
		category := domain.Category{}

		category.Name = req.Name
		category.CreatedAt = time.Now().UnixMilli()

		categories = append(categories, category)
	}

	service.CategoryRepo.SaveMany(ctx, tx, categories)

	return helper.ToCategoryResponses(categories)
}

func (service *CategoryServiceImpl) DeleteAll(ctx context.Context) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.CategoryRepo.DeleteAll(ctx, tx)
}
