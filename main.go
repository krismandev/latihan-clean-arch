package main

import (
	"agit-test/app"
	"agit-test/controller"
	"agit-test/middleware"
	"agit-test/repository"
	"agit-test/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()
	validate := validator.New()
	db := app.NewDB()

	karyawanRepository := repository.NewKaryawanRepository()
	karyawanService := service.NewKaryawanService(karyawanRepository, db, validate)
	karyawanController := controller.NewKaryawanController(karyawanService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	karyawanRouter := router.Group("karyawan")
	{
		karyawanRouter.Use(middleware.Authentication())
		karyawanRouter.GET("/", karyawanController.FindAll)
		karyawanRouter.GET("/:karyawanId", karyawanController.FindById)
		karyawanRouter.POST("/", karyawanController.Create)
		karyawanRouter.PUT("/:karyawanId", karyawanController.Update)
		karyawanRouter.DELETE("/:karyawanId", karyawanController.Delete)
	}
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	router.Run(":4000")
}
