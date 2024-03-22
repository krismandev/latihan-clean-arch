package service

import (
	"agit-test/model/web"
	"context"
)

// untuk response dibedakan dari object domain nya. tujaunnya agar kita dapat mengatur/ membatasi properti apa saja yang boleh dikembalikan sbg response
type KaryawanService interface {
	Create(ctx context.Context, request web.KaryawanCreateRequest) (web.KaryawanResponse, error)
	Update(ctx context.Context, requst web.KaryawanUpdateRequest, karyawanId int) (web.KaryawanResponse, error)
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.KaryawanResponse
	FindAll(ctx context.Context) []web.KaryawanResponse
}
