package repository

import (
	"context"
	"database/sql"

	"github.com/Jehanv60/model/domain"
)

type PenggunaRepository interface {
	Save(ctx context.Context, tx *sql.Tx, pengguna domain.Pengguna) domain.Pengguna
	Update(ctx context.Context, tx *sql.Tx, pengguna domain.Pengguna) domain.Pengguna
	FindById(ctx context.Context, tx *sql.Tx, penggunaId int) domain.Pengguna
	FindByPenggunaRegister(ctx context.Context, tx *sql.Tx, NamaPengguna, Email string) domain.Pengguna
	FindByPenggunaLogin(ctx context.Context, tx *sql.Tx, NamaPengguna string) domain.Pengguna
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Pengguna
	Delete(ctx context.Context, tx *sql.Tx, pengguna domain.Pengguna)
}
