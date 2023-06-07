package service

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/faridlan/nostra-api-product/exception"
	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/helper/hash"
	"github.com/faridlan/nostra-api-product/helper/mysql"
	"github.com/faridlan/nostra-api-product/model/domain"
	"github.com/faridlan/nostra-api-product/model/web"
	"github.com/faridlan/nostra-api-product/repository"
	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	UserRepo repository.UserRepository
	RoleRepo repository.RoleRepository
	DB       *sql.DB
	Validate *validator.Validate
}

func NewAuthService(userRepo repository.UserRepository, roleRepo repository.RoleRepository, db *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepo: userRepo,
		RoleRepo: roleRepo,
		DB:       db,
		Validate: validate,
	}
}

func (service *AuthServiceImpl) Register(ctx context.Context, request web.UserCreateReq) web.LoginResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	errors := helper.TranslateError(err, service.Validate)
	if err != nil {
		panic(exception.NewValidationError(errors))
	}

	_, err = service.UserRepo.FindName(ctx, tx, request.Username)
	if err != nil {
		panic(exception.NewInterfaceErrorUnauth(err.Error()))
	}

	imageString := mysql.NewNullString(request.Image)
	hash := hash.HashAndSalt([]byte(request.Password))

	user := domain.User{
		UserId:   0,
		Id:       request.Id,
		Username: request.Username,
		Password: hash,
		Email:    request.Email,
		Image:    imageString,
		Role: domain.Role{
			Id: request.RoleId,
		},
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: &mysql.NullInt{},
	}

	role, err := service.RoleRepo.FindByName(ctx, tx, "user")
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

	user.Role.Id = role.Id

	user = service.UserRepo.Save(ctx, tx, user)
	user, err = service.UserRepo.FindId(ctx, tx, user.UserId)
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

	tokenString := helper.JwtGen(user)
	userResponseLogin := helper.ToLoginResponse(user)
	userResponseLogin.Token = tokenString

	return userResponseLogin

}

func (service *AuthServiceImpl) Update(ctx context.Context, request web.UserUpdateReq) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	errors := helper.TranslateError(err, service.Validate)
	if err != nil {
		panic(exception.NewValidationError(errors))
	}

	user, err := service.UserRepo.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

	imageString := mysql.NewNullString(request.Image)
	updateInt := mysql.NewNullInt64(time.Now().UnixMilli())

	user.Username = request.Username
	user.Email = request.Email
	user.Image = imageString
	user.Role.Id = request.RoleId
	user.UpdatedAt = updateInt

	user = service.UserRepo.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *AuthServiceImpl) FindById(ctx context.Context, userId string) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepo.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewInterfaceError(err.Error()))
	}

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

func (service *AuthServiceImpl) Login(ctx context.Context, request web.Login) web.LoginResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	errors := helper.TranslateError(err, service.Validate)
	if err != nil {
		panic(exception.NewValidationError(errors))
	}

	userReq := domain.User{
		Username: request.Username,
	}

	UserResponse, _ := service.UserRepo.Login(ctx, tx, userReq)
	log.Println(UserResponse)

	err = hash.ComparePassword(UserResponse.Password, []byte(request.Password))
	if err != nil {
		panic(exception.NewInterfaceErrorUnauth(err.Error()))
	}

	tokenString := helper.JwtGen(UserResponse)
	userResponseLogin := helper.ToLoginResponse(UserResponse)
	userResponseLogin.Token = tokenString

	return userResponseLogin
}

func (service *AuthServiceImpl) SaveMany(ctx context.Context, request []web.UserCreateReq) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := []domain.User{}

	for _, req := range request {
		user := domain.User{}

		imageString := mysql.NewNullString(req.Image)
		hash := hash.HashAndSalt([]byte(req.Password))

		user.Username = req.Username
		user.Password = hash
		user.Email = req.Email
		user.Image = imageString
		user.Role.Id = req.RoleId
		user.CreatedAt = time.Now().UnixMilli()

		users = append(users, user)
	}

	service.UserRepo.SaveMany(ctx, tx, users)
	users = service.UserRepo.FindAll(ctx, tx)

	return helper.ToUserResponses(users)
}

func (service *AuthServiceImpl) DeleteAll(ctx context.Context) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.UserRepo.DeleteAll(ctx, tx)
}
