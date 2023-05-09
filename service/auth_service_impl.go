package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/helper/mysql"
	"github.com/faridlan/nostra-api-product/model/domain"
	"github.com/faridlan/nostra-api-product/model/web"
	"github.com/faridlan/nostra-api-product/repository"
)

type AuthServiceImpl struct {
	UserRepo repository.UserRepository
	DB       *sql.DB
}

func NewAuthService(userRepo repository.UserRepository, db *sql.DB) AuthService {
	return &AuthServiceImpl{
		UserRepo: userRepo,
		DB:       db,
	}
}

func (service *AuthServiceImpl) Register(ctx context.Context, request web.UserCreateReq) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	imageString := mysql.NewNullString(request.Image)
	user := domain.User{
		Id:        request.Id,
		Username:  request.Username,
		Password:  request.Password,
		Email:     request.Email,
		Image:     imageString,
		RoleId:    request.RoleId,
		CreatedAt: time.Now().UnixMilli(),
	}

	user = service.UserRepo.Save(ctx, tx, user)
	user, err = service.UserRepo.FindId(ctx, tx, user.UserId)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

func (service *AuthServiceImpl) Update(ctx context.Context, request web.UserUpdateReq) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepo.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	imageString := mysql.NewNullString(request.Image)
	updateInt := mysql.NewNullInt64(time.Now().UnixMilli())

	user.Username = request.Username
	user.Email = request.Email
	user.Image = imageString
	user.RoleId = request.RoleId
	user.UpdatedAt = updateInt

	user = service.UserRepo.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *AuthServiceImpl) FindById(ctx context.Context, userId string) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepo.FindById(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

func (service *AuthServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepo.FindAll(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToUserResponses(users)
}

func (service *AuthServiceImpl) SaveMany(ctx context.Context, request []web.UserCreateReq) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := []domain.User{}

	for _, req := range request {
		user := domain.User{}

		imageString := mysql.NewNullString(req.Image)

		user.Username = req.Username
		user.Password = req.Password
		user.Email = req.Email
		user.Image = imageString
		user.RoleId = req.RoleId
		user.CreatedAt = time.Now().UnixMilli()

		users = append(users, user)
	}

	users = service.UserRepo.SaveMany(ctx, tx, users)

	return helper.ToUserResponses(users)
}

func (service *AuthServiceImpl) DeleteAll(ctx context.Context) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.UserRepo.DeleteAll(ctx, tx)
}
