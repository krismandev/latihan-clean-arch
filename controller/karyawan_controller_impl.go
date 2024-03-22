package controller

import (
	"agit-test/helper"
	"agit-test/model/web"
	"agit-test/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KaryawanControllerImpl struct {
	KaryawanService service.KaryawanService
}

func NewKaryawanController(karyawanService service.KaryawanService) KaryawanController {
	return &KaryawanControllerImpl{
		KaryawanService: karyawanService,
	}
}

func (controller *KaryawanControllerImpl) Create(c *gin.Context) {
	var err error

	ctx := c.Request.Context()

	karyawanCreateRequest := web.KaryawanCreateRequest{}

	helper.ReadFromJSON(c, &karyawanCreateRequest)
	karyawanCreateResponse, err := controller.KaryawanService.Create(ctx, karyawanCreateRequest)
	helper.WriteResponseJSON(c, karyawanCreateResponse, err)
}

func (controller *KaryawanControllerImpl) Update(c *gin.Context) {
	var err error
	karyawanUpdateRequest := web.KaryawanUpdateRequest{}

	var karyawanUpdateResponse web.KaryawanResponse

	err = c.Bind(&karyawanUpdateRequest)
	helper.PanicIfError(err)
	karyawanId := c.Param("karyawanId")
	karyawanIdInt, err := strconv.Atoi(karyawanId)
	if err != nil {
		helper.WriteResponseJSON(c, karyawanUpdateResponse, err)
	}

	karyawanCreateResponse, err := controller.KaryawanService.Update(c.Request.Context(), karyawanUpdateRequest, karyawanIdInt)

	helper.WriteResponseJSON(c, karyawanCreateResponse, err)
}

func (controller *KaryawanControllerImpl) Delete(c *gin.Context) {
	var err error
	karyawanId := c.Param("karyawanId")

	id, err := strconv.Atoi(karyawanId)
	helper.PanicIfError(err)

	controller.KaryawanService.Delete(c.Request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteResponseJSON(c, webResponse, err)
	// helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KaryawanControllerImpl) FindById(c *gin.Context) {
	var karyawanResponse interface{}

	karyawanId := c.Param("karyawanId")

	id, err := strconv.Atoi(karyawanId)
	helper.PanicIfError(err)

	karyawan := controller.KaryawanService.FindById(c.Request.Context(), id)
	if karyawan.Id != 0 {
		karyawanResponse = karyawan
	}

	helper.WriteResponseJSON(c, karyawanResponse, err)
}

func (controller *KaryawanControllerImpl) FindAll(c *gin.Context) {
	var err error
	ctx := c.Request.Context()
	karyawanResponses := controller.KaryawanService.FindAll(ctx)

	helper.WriteResponseJSON(c, karyawanResponses, err)
}
