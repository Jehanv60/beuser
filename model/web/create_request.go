package web

type PenggunaCreateRequest struct {
	Pengguna string `validate:"required,max=100,min=1,alphanum" json:"pengguna"`
	Email    string `validate:"required,email" json:"email"`
	Sandi    string `validate:"required,max=100,min=1,alphanum" json:"sandi"`
}
