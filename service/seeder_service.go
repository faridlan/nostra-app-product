package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/helper/seeder"
	"github.com/faridlan/nostra-api-product/model/web"
	"github.com/faridlan/nostra-api-product/repository"
)

type SeederService interface {
	SaveMany(ctx context.Context, role []web.RoleCreateReq, user []web.UserCreateReq, category []web.CategoryCreateReq, product []web.ProductCreateReq)
	Delete(ctx context.Context)
}

type SeederServiceImpl struct {
	DB           *sql.DB
	RoleRepo     repository.RoleRepository
	UserRepo     repository.UserRepository
	CategoryRepo repository.CategoryRepository
	ProductRepo  repository.ProductRepository
}

func NewSeederService(db *sql.DB, roleRepo repository.RoleRepository,
	UserRepo repository.UserRepository,
	CategoryRepo repository.CategoryRepository,
	ProductRepo repository.ProductRepository,
) SeederService {
	return &SeederServiceImpl{
		DB:           db,
		RoleRepo:     roleRepo,
		UserRepo:     UserRepo,
		CategoryRepo: CategoryRepo,
		ProductRepo:  ProductRepo,
	}
}

func (service *SeederServiceImpl) Delete(ctx context.Context) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.ProductRepo.DeleteImage(ctx, tx)
	service.UserRepo.DeleteAll(ctx, tx)
	service.RoleRepo.DeleteAll(ctx, tx)
	service.ProductRepo.DeleteAll(ctx, tx)
	service.CategoryRepo.DeleteAll(ctx, tx)
}

func (service *SeederServiceImpl) SaveMany(ctx context.Context, roles []web.RoleCreateReq, users []web.UserCreateReq, categories []web.CategoryCreateReq, products []web.ProductCreateReq) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	domainUser := []web.UserCreateReq{}

	rolesResponse := service.RoleRepo.SaveMany(ctx, tx, seeder.RoleCreateMany(roles))

	for i, user := range users {
		roles, err := service.RoleRepo.FindByName(ctx, tx, rolesResponse[i].Name)
		helper.PanicIfError(err)
		user.RoleId = roles.RoleId
		domainUser = append(domainUser, user)
	}

	service.UserRepo.SaveMany(ctx, tx, seeder.UserCreateMany(domainUser))

	domainProduct := []web.ProductCreateReq{}

	categoryResponse := service.CategoryRepo.SaveMany(ctx, tx, seeder.CategoryCreateMany(categories))

	for i, product := range products {
		category, err := service.CategoryRepo.FindByName(ctx, tx, categoryResponse[i].Name)
		helper.PanicIfError(err)
		product.CategoryId = category.CategoryId
		domainProduct = append(domainProduct, product)
	}

	productsResult := service.ProductRepo.SaveMany(ctx, tx, seeder.ProductCreateMany(domainProduct))

	// fmt.Println(productsResult)

	service.ProductRepo.SaveImage(ctx, tx, seeder.ProductImageCreate(productsResult))

}
