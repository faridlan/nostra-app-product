package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/nostra-api-product/model/domain"
)

type UserRepository interface {
	SaveMany(ctx context.Context, tx *sql.Tx, user []domain.User) []domain.User
	DeleteAll(ctx context.Context, tx *sql.Tx)
}
