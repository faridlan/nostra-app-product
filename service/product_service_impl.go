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

type ProductServiceImpl struct {
	ProductRepo repository.ProductRepository
	DB          *sql.DB
	Validate    *validator.Validate
}

func NewProductService(productRepo repository.ProductRepository, db *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepo: productRepo,
		DB:          db,
		Validate:    validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateReq) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	errors := helper.TranslateError(err, service.Validate)
	if err != nil {
		panic(exception.NewValidationError(errors))
	}

	imageString := mysql.NewNullString(request.Image)

	product := domain.Product{
		Name:        request.Name,
		Price:       request.Price,
		Quantity:    request.Quantity,
		Description: request.Description,
		Image:       imageString,
		Category: domain.Category{
			Id: request.CategoryId,
		},
		CreatedAt: time.Now().UnixMilli(),
	}

	productResult := service.ProductRepo.Save(ctx, tx, product)
	productResult, err = service.ProductRepo.FindId(ctx, tx, productResult.ProductId)
	helper.PanicIfError(err)

	return helper.ToProductResponse(productResult)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateReq) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	errors := helper.TranslateError(err, service.Validate)
	if err != nil {
		panic(exception.NewValidationError(errors))
	}

	imageString := mysql.NewNullString(request.Image)
	upddateInt := mysql.NewNullInt64(time.Now().UnixMilli())

	product, err := service.ProductRepo.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

	product.Name = request.Name
	product.Price = request.Price
	product.Quantity = request.Quantity
	product.Description = request.Description
	product.Image = imageString
	product.Category.Id = request.CategoryId
	product.UpdatedAt = upddateInt

	productResult := service.ProductRepo.Update(ctx, tx, product)

	return helper.ToProductResponse(productResult)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepo.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

	service.ProductRepo.Delete(ctx, tx, product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId string) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepo.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepo.FindAll(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToProductResponses(products)
}

func (service *ProductServiceImpl) CreateMany(ctx context.Context, request []web.ProductCreateReq) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := []domain.Product{}

	for _, req := range request {
		product := domain.Product{}
		imageString := mysql.NewNullString(req.Image)

		product.Name = req.Name
		product.Price = req.Price
		product.Quantity = req.Quantity
		product.Description = req.Description
		product.Image = imageString
		product.Category.Id = req.CategoryId
		product.CreatedAt = time.Now().UnixMilli()

		products = append(products, product)
	}

	service.ProductRepo.SaveMany(ctx, tx, products)
	productResponses := service.ProductRepo.FindAll(ctx, tx)

	return helper.ToProductResponses(productResponses)
}

func (service *ProductServiceImpl) DeleteAll(ctx context.Context) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.ProductRepo.DeleteAll(ctx, tx)
}
