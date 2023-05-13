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

type RoleServiceImpl struct {
	RoleRepo repository.RoleRepository
	DB       *sql.DB
	Validate *validator.Validate
}

func NewRoleService(roleRepo repository.RoleRepository, db *sql.DB, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepo: roleRepo,
		DB:       db,
		Validate: validate,
	}
}

func (service *RoleServiceImpl) Create(ctx context.Context, request web.RoleCreateReq) web.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	errors := helper.TranslateError(err, service.Validate)

	if err != nil {
		panic(exception.NewValidationError(errors))
	}

	role := domain.Role{
		Name:      request.Name,
		CreatedAt: time.Now().UnixMilli(),
	}

	role = service.RoleRepo.Save(ctx, tx, role)

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) Update(ctx context.Context, request web.RoleUpdateReq) web.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	errors := helper.TranslateError(err, service.Validate)

	if err != nil {
		panic(exception.NewValidationError(errors))
	}

	role, err := service.RoleRepo.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}
	roleUpdate := mysql.NewNullInt64(time.Now().UnixMilli())

	role.Name = request.Name
	role.UpdatedAt = roleUpdate

	role = service.RoleRepo.Update(ctx, tx, role)

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) FindById(ctx context.Context, roleId string) web.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role, err := service.RoleRepo.FindById(ctx, tx, roleId)
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) FindAll(ctx context.Context) []web.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roles := service.RoleRepo.FindAll(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToRoleResponses(roles)
}

func (service *RoleServiceImpl) CreateMany(ctx context.Context, requests []web.RoleCreateReq) []web.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roles := []domain.Role{}

	for _, role := range requests {
		roleStruct := domain.Role{}
		roleStruct.Name = role.Name
		roleStruct.CreatedAt = time.Now().UnixMilli()
		roles = append(roles, roleStruct)
	}

	roles = service.RoleRepo.SaveMany(ctx, tx, roles)

	return helper.ToRoleResponses(roles)

}

func (service *RoleServiceImpl) DeleteAll(ctx context.Context) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.RoleRepo.DeleteAll(ctx, tx)
}
