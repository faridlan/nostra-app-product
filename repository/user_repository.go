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
	FindName(ctx context.Context, tx *sql.Tx, username string) (string, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	SaveMany(ctx context.Context, tx *sql.Tx, user []domain.User) []domain.User
	DeleteAll(ctx context.Context, tx *sql.Tx)
	SaveWL(ctx context.Context, tx *sql.Tx, whitelist domain.Whitelist) domain.Whitelist
	DeleteWL(ctx context.Context, tx *sql.Tx, userId string)
}
