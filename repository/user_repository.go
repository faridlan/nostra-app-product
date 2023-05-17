package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/nostra-api-product/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId string) (domain.User, error)
	FindId(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	SaveMany(ctx context.Context, tx *sql.Tx, user []domain.User) []domain.User
	DeleteAll(ctx context.Context, tx *sql.Tx)
}
