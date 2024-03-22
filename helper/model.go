package helper

import (
	"agit-test/model/domain"
	"agit-test/model/web"
)

func ToKaryawanResponse(karyawan domain.Karyawan) web.KaryawanResponse {
	var karyawanResponse web.KaryawanResponse
	karyawanResponse.Id = karyawan.Id
	karyawanResponse.Nama = karyawan.Nama
	karyawanResponse.Nip = karyawan.Nip
	karyawanResponse.TempatLahir = karyawan.TempatLahir
	karyawanResponse.TanggalLahir = karyawan.TanggalLahir
	karyawanResponse.Agama = karyawan.Agama
	karyawanResponse.Alamat = karyawan.Alamat
	karyawanResponse.JenisKelamin = karyawan.JenisKelamin
	karyawanResponse.Umur = karyawan.Umur
	karyawanResponse.Email = karyawan.Email
	karyawanResponse.NoHandphone = karyawan.NoHandphone
	karyawanResponse.CreatedAt = karyawan.CreatedAt
	karyawanResponse.UpdatedAt = karyawan.UpdatedAt
	karyawanResponse.DeletedAt = karyawan.DeletedAt
	return karyawanResponse
}

func ToUserResponse(user domain.User) web.UserResponse {
	var userResponse web.UserResponse
	userResponse.Username = user.Username
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt
	userResponse.DeletedAt = user.DeletedAt

	return userResponse
}
