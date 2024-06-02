package service

import (
	"context"

	"github.com/Jehanv60/model/web"
)

type PenggunaService interface {
	Create(ctx context.Context, request web.PenggunaCreateRequest) web.PenggunaResponse
	Update(ctx context.Context, update web.PenggunaUpdate) web.PenggunaResponse
	FindById(ctx context.Context, penggunaId int) web.PenggunaResponse
	FindByPenggunaLogin(ctx context.Context, NamaPengguna string) web.PenggunaResponse
	FindAll(ctx context.Context) []web.PenggunaResponse
	Delete(ctx context.Context, barangId int)
}
