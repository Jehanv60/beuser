package web

type PenggunaResponse struct {
	Id       int    `json:"id"`
	Pengguna string `json:"pengguna"`
	Email    string `json:"email"`
	Sandi    string `json:"sandi"`
}
