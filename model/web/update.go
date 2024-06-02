package web

type PenggunaUpdate struct {
	Id       int    `validate:"required"`
	Pengguna string `validate:"required,max=100,min=1,alphanumdash" json:"pengguna"`
	Email    string `validate:"required,email" json:"email"`
	Sandi    string `validate:"required,max=100,min=1,alphanumdash" json:"sandi"`
}
