package service

import (
	"agit-test/helper"
	"agit-test/model/domain"
	"agit-test/model/web"
	"agit-test/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

//database transaction di innitiate di level service karna kemungkinan dalam suatu transaksi bisa menggunakan lebih dari 1 repository.
//cth repository oder dan repository order detail

type KaryawanServiceImpl struct {
	KaryawanRepository repository.KaryawanRepository
	// karna DB ini asalnya adalah struct. bukan interface. jadi kita pakai pointer.
	// kalau bentuknya merupakan interface, kita tidak pakai pointer.
	DB       *sql.DB
	Validate *validator.Validate
}

func NewKaryawanService(karyawanRepository repository.KaryawanRepository, DB *sql.DB, validate *validator.Validate) KaryawanService {
	return &KaryawanServiceImpl{
		KaryawanRepository: karyawanRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *KaryawanServiceImpl) Create(ctx context.Context, request web.KaryawanCreateRequest) (web.KaryawanResponse, error) {
	var err error
	var resp web.KaryawanResponse
	err = service.Validate.Struct(request)
	if err != nil {
		return resp, err
	}
	tx, err := service.DB.Begin()
	if err != nil {
		return resp, err
	}

	defer helper.CommitOrRollback(tx)

	karyawan := domain.Karyawan{
		Nama:         request.Nama,
		Nip:          request.Nip,
		TempatLahir:  request.TempatLahir,
		TanggalLahir: request.TanggalLahir,
		Umur:         request.Umur,
		Alamat:       request.Alamat,
		Agama:        request.Agama,
		JenisKelamin: request.JenisKelamin,
		NoHandphone:  request.NoHandphone,
		Email:        request.Email,
	}

	karyawan = service.KaryawanRepository.Save(ctx, tx, karyawan)

	resp = helper.ToKaryawanResponse(karyawan)
	return resp, err

}

func (service *KaryawanServiceImpl) Update(ctx context.Context, request web.KaryawanUpdateRequest, karyawanId int) (web.KaryawanResponse, error) {
	var err error
	var resp web.KaryawanResponse
	err = service.Validate.Struct(request)
	if err != nil {
		return resp, err
	}
	tx, err := service.DB.Begin()
	if err != nil {
		return resp, err
	}
	defer helper.CommitOrRollback(tx)

	karyawan := service.KaryawanRepository.FindById(ctx, tx, karyawanId)

	if err != nil {
		return resp, err
	}

	karyawan.Nama = request.Nama
	karyawan.Nip = request.Nip
	karyawan.Agama = request.Agama
	karyawan.Alamat = request.Alamat
	karyawan.JenisKelamin = request.JenisKelamin
	karyawan.Email = request.Email
	karyawan.TempatLahir = request.TempatLahir
	karyawan.TanggalLahir = request.TanggalLahir
	karyawan.NoHandphone = request.NoHandphone
	karyawan.Umur = request.Umur

	karyawan = service.KaryawanRepository.Update(ctx, tx, karyawan, karyawanId)
	resp = helper.ToKaryawanResponse(karyawan)

	return resp, err
}

func (service *KaryawanServiceImpl) Delete(ctx context.Context, karyawanId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	karyawan := service.KaryawanRepository.FindById(ctx, tx, karyawanId)

	service.KaryawanRepository.Delete(ctx, tx, karyawan)

}

func (service *KaryawanServiceImpl) FindById(ctx context.Context, karyawanId int) web.KaryawanResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	karyawan := service.KaryawanRepository.FindById(ctx, tx, karyawanId)

	return helper.ToKaryawanResponse(karyawan)
}

func (service *KaryawanServiceImpl) FindAll(ctx context.Context) []web.KaryawanResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	karyawanList := service.KaryawanRepository.FindAll(ctx, tx)

	var responseData []web.KaryawanResponse
	for _, each := range karyawanList {
		dt := helper.ToKaryawanResponse(each)
		responseData = append(responseData, dt)
	}

	return responseData
}
