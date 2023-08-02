package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/nostra-api-product/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, product domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, productId string) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
	FindId(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error)
	FindProductImages(ctx context.Context, tx *sql.Tx, productId int) []domain.ProductImage
	SaveImage(ctx context.Context, tx *sql.Tx, products []domain.ProductImage) []domain.ProductImage
	DeleteImage(ctx context.Context, tx *sql.Tx)
	SaveMany(ctx context.Context, tx *sql.Tx, products []domain.Product) []domain.Product
	DeleteAll(ctx context.Context, tx *sql.Tx)
}
