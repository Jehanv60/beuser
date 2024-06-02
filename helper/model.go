package helper

import (
	"github.com/Jehanv60/model/domain"
	"github.com/Jehanv60/model/web"
)

func ToPenggunaResponse(pengguna domain.Pengguna) web.PenggunaResponse {
	return web.PenggunaResponse{
		Id:       pengguna.Id,
		Pengguna: pengguna.Pengguna,
		Email:    pengguna.Email,
		Sandi:    pengguna.Sandi,
	}
}

func ToPenggunaResponses(penggunas []domain.Pengguna) []web.PenggunaResponse {
	var penggunaResponses []web.PenggunaResponse
	for _, penggunass := range penggunas {
		penggunaResponses = append(penggunaResponses, ToPenggunaResponse(penggunass))
	}
	return penggunaResponses
}
