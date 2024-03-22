package service

import (
	"agit-test/model/web"
	"context"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error)
	Login(ctx context.Context, request web.LoginRequest) (web.LoginResponse, error)
}
