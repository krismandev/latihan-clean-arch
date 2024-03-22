package controller

import (
	"agit-test/helper"
	"agit-test/model/web"
	"agit-test/service"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(c *gin.Context) {
	var err error
	ctx := c.Request.Context()
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromJSON(c, &userCreateRequest)
	karyawanCreateResponse, err := controller.UserService.Create(ctx, userCreateRequest)
	helper.WriteResponseJSON(c, karyawanCreateResponse, err)
}

func (controller UserControllerImpl) Login(c *gin.Context) {
	var err error
	ctx := c.Request.Context()
	loginRequest := web.LoginRequest{}
	helper.ReadFromJSON(c, &loginRequest)
	loginResponse, err := controller.UserService.Login(ctx, loginRequest)
	helper.WriteResponseJSON(c, loginResponse, err)
}
