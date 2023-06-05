package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/faridlan/nostra-api-product/exception"
	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/helper/mysql"
	"github.com/faridlan/nostra-api-product/model/domain"
	"github.com/faridlan/nostra-api-product/model/web"
	"github.com/faridlan/nostra-api-product/repository"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepo repository.CategoryRepository
	DB           *sql.DB
	Validate     *validator.Validate
}

func NewCategoryService(categoryRepo repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepo: categoryRepo,
		DB:           db,
		Validate:     validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateReq) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	errors := helper.TranslateError(err, service.Validate)
	if err != nil {
		panic(exception.NewValidationError(errors))
	}

	cateogry := domain.Category{
		Name:      request.Name,
		CreatedAt: time.Now().UnixMilli(),
	}

	categoryResult := service.CategoryRepo.Save(ctx, tx, cateogry)

	categoryResult, err = service.CategoryRepo.FindId(ctx, tx, categoryResult.CateogoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(categoryResult)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateReq) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	errors := helper.TranslateError(err, service.Validate)
	if err != nil {
		panic(exception.NewValidationError(errors))
	}

	updateAt := mysql.NewNullInt64(time.Now().UnixMilli())

	category, err := service.CategoryRepo.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

	category.Name = request.Name
	category.UpdatedAt = updateAt

	categoryResult := service.CategoryRepo.Update(ctx, tx, category)

	// categoryResult, err = service.CategoryRepo.FindId(ctx, tx, categoryResult.CateogoryId)
	// helper.PanicIfError(err)

	return helper.ToCategoryResponse(categoryResult)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepo.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepo.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId string) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepo.FindById(ctx, tx, categoryId)

	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepo.FindAll(ctx, tx)
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

	return helper.ToCategoryResponses(categories)
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
