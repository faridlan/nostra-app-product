package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/nostra-api-product/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId string) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	FindId(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindByName(ctx context.Context, tx *sql.Tx, categoryName string) (domain.Category, error)
	SaveMany(ctx context.Context, tx *sql.Tx, categories []domain.Category) []domain.Category
	DeleteAll(ctx context.Context, tx *sql.Tx)
}
