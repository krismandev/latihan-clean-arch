package repository

import (
	"agit-test/helper"
	"agit-test/model/domain"
	"context"
	"database/sql"
	"strconv"
	"time"
)

type KaryawanRepositoryImpl struct {
}

func NewKaryawanRepository() KaryawanRepository {
	return &KaryawanRepositoryImpl{}
}

func (repository *KaryawanRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) domain.Karyawan {
	var baseIn string
	var baseParam []interface{}

	helper.AppendComma(&baseIn, &baseParam, "?", karyawan.Nama)
	helper.AppendComma(&baseIn, &baseParam, "?", karyawan.Nip)
	helper.AppendComma(&baseIn, &baseParam, "?", karyawan.TempatLahir)
	helper.AppendComma(&baseIn, &baseParam, "?", karyawan.TanggalLahir)
	helper.AppendComma(&baseIn, &baseParam, "?", strconv.Itoa(karyawan.Umur))
	helper.AppendComma(&baseIn, &baseParam, "?", karyawan.Alamat)
	helper.AppendComma(&baseIn, &baseParam, "?", karyawan.Agama)
	helper.AppendComma(&baseIn, &baseParam, "?", karyawan.JenisKelamin)
	helper.AppendComma(&baseIn, &baseParam, "?", karyawan.NoHandphone)
	helper.AppendComma(&baseIn, &baseParam, "?", karyawan.Email)
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	helper.AppendComma(&baseIn, &baseParam, "?", timeNow)

	query := "INSERT INTO karyawan (nama,nip,tempat_lahir,tanggal_lahir,umur,alamat,agama,jenis_kelamin,no_handphone,email,created_at) VALUES (" + baseIn + ")"
	res, err := tx.ExecContext(ctx, query, baseParam...)
	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	karyawan.Id = int(id)
	karyawan.CreatedAt = timeNow
	return karyawan
}

func (repository *KaryawanRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan, karyawanId int) domain.Karyawan {
	var baseUp string
	var baseParam []interface{}

	helper.AppendComma(&baseUp, &baseParam, "nama = ?", karyawan.Nama)
	helper.AppendComma(&baseUp, &baseParam, "nip = ?", karyawan.Nip)
	helper.AppendComma(&baseUp, &baseParam, "tempat_lahir = ?", karyawan.TempatLahir)
	helper.AppendComma(&baseUp, &baseParam, "tanggal_lahir = ?", karyawan.TanggalLahir)
	helper.AppendComma(&baseUp, &baseParam, "umur = ?", strconv.Itoa(karyawan.Umur))
	helper.AppendComma(&baseUp, &baseParam, "alamat = ?", karyawan.Alamat)
	helper.AppendComma(&baseUp, &baseParam, "agama = ?", karyawan.Agama)
	helper.AppendComma(&baseUp, &baseParam, "jenis_kelamin = ?", karyawan.JenisKelamin)
	helper.AppendComma(&baseUp, &baseParam, "no_handphone = ?", karyawan.NoHandphone)
	helper.AppendComma(&baseUp, &baseParam, "email = ?", karyawan.Email)
	helper.AppendCommaRaw(&baseUp, " updated_at =  now()")

	query := "UPDATE karyawan set " + baseUp + " WHERE id = ?"
	baseParam = append(baseParam, karyawanId)
	_, err := tx.ExecContext(ctx, query, baseParam...)
	helper.PanicIfError(err)

	return karyawan
}

func (repository *KaryawanRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) {
	query := "UPDATE karyawan set deleted_at = now() where id  = ?"
	_, err := tx.ExecContext(ctx, query, karyawan.Id)
	helper.PanicIfError(err)
}

func (repository *KaryawanRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, karyawanId int) domain.Karyawan {
	query := "SELECT id, nama, nip, tempat_lahir, tanggal_lahir, umur, alamat, agama, jenis_kelamin, no_handphone, email, created_at, updated_at, deleted_at FROM karyawan WHERE id = ? AND deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, query, karyawanId)
	defer rows.Close()
	helper.PanicIfError(err)
	var karyawan domain.Karyawan
	if rows.Next() {
		rows.Scan(&karyawan.Id, &karyawan.Nama, &karyawan.Nip, &karyawan.TempatLahir, &karyawan.TanggalLahir, &karyawan.Umur, &karyawan.Alamat, &karyawan.Agama, &karyawan.JenisKelamin, &karyawan.NoHandphone, &karyawan.Email, &karyawan.CreatedAt, &karyawan.UpdatedAt, &karyawan.DeletedAt)
	}

	return karyawan
}

func (repository *KaryawanRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Karyawan {
	query := "SELECT id, nama, nip, tempat_lahir, tanggal_lahir, umur, alamat, agama, jenis_kelamin, no_handphone, email, created_at, updated_at, deleted_at FROM karyawan WHERE deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, query)
	defer rows.Close()
	helper.PanicIfError(err)
	var listKaryawan []domain.Karyawan
	for rows.Next() {

		var karyawan domain.Karyawan
		rows.Scan(&karyawan.Id, &karyawan.Nama, &karyawan.Nip, &karyawan.TempatLahir, &karyawan.TanggalLahir, &karyawan.Umur, &karyawan.Alamat, &karyawan.Agama, &karyawan.JenisKelamin, &karyawan.NoHandphone, &karyawan.Email, &karyawan.CreatedAt, &karyawan.UpdatedAt, &karyawan.DeletedAt)

		listKaryawan = append(listKaryawan, karyawan)
	}

	return listKaryawan
}
