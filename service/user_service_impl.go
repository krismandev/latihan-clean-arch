package service

import (
	"agit-test/helper"
	"agit-test/model/domain"
	"agit-test/model/web"
	"agit-test/repository"
	"context"
	"database/sql"

	"errors"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	// karna DB ini asalnya adalah struct. bukan interface. jadi kita pakai pointer.
	// kalau bentuknya merupakan interface, kita tidak pakai pointer.
	DB       *sql.DB
	Validate *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error) {
	var err error
	var response web.UserResponse

	err = service.Validate.Struct(request)
	if err != nil {
		return response, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return response, err
	}

	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Username: request.Username,
		Password: helper.HashPassword(request.Password),
	}

	user = service.UserRepository.Save(ctx, tx, user)

	response = helper.ToUserResponse(user)

	return response, err
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.LoginRequest) (web.LoginResponse, error) {
	var err error
	var response web.LoginResponse
	err = service.Validate.Struct(request)
	if err != nil {
		return response, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return response, err
	}

	defer helper.CommitOrRollback(tx)

	user := service.UserRepository.FindByUsername(ctx, tx, request.Username)

	checkPassword := helper.ComparePass([]byte(user.Password), []byte(request.Password))
	if !checkPassword {
		return response, errors.New("invalid username or password")
	}

	token := helper.GenerateToken(uint(user.Id), user.Username)

	response.Id = user.Id
	response.Username = user.Username
	response.Token = token

	return response, err
}
