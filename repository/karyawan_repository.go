package repository

import (
	"agit-test/model/domain"
	"context"
	"database/sql"
)

type KaryawanRepository interface {
	Save(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan) domain.Karyawan
	Update(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan, karyawanId int) domain.Karyawan
	Delete(ctx context.Context, tx *sql.Tx, karyawan domain.Karyawan)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Karyawan
	FindById(ctx context.Context, tx *sql.Tx, id int) domain.Karyawan
}
