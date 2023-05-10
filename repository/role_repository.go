package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/nostra-api-product/model/domain"
)

type RoleRepository interface {
	Save(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role
	Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role
	FindById(ctx context.Context, tx *sql.Tx, roleId string) (domain.Role, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Role
	SaveMany(ctx context.Context, tx *sql.Tx, roles []domain.Role) []domain.Role
	DeleteAll(ctx context.Context, tx *sql.Tx)
}
