package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO products(product_id, name, price, quantity, description, image, category_id, created_at) values (UUID_TO_BIN(UUID()),?,?,?,?,?,UUID_TO_BIN(?),?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Quantity, product.Description, product.Image, product.CategoryId, product.CreatedAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.ProductId = int(id)

	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE products SET name = ?, price = ?, quantity = ?, description = ?, image = ?, category_id = UUID_TO_BIN(?), updated_at = ? WHERE REPLACE(BIN_TO_UUID(product_id), '-', '') = ?"
	_, err := tx.ExecContext(ctx, SQL, ctx, product.Name, product.Price, product.Quantity, product.Description, product.Image, product.CategoryId, product.UpdatedAt, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "DELETE products WHERE REPLACE(BIN_TO_UUID(product_id), '-', '') = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}
func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId string) (domain.Product, error) {
	SQL := `SELECT REPLACE(BIN_TO_UUID(product_id), '-', ''), name, price, quantity, description, image, REPLACE(BIN_TO_UUID(category_id), '-', ''), 
	created_at, updated_at FROM products WHERE REPLACE(BIN_TO_UUID(product_id), '-', '') = ?`

	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)

	defer rows.Close()

	product := domain.Product{}

	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.Image, &product.CategoryId, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicIfError(err)

		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}

func (repository *ProductRepositoryImpl) FindId(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := `SELECT REPLACE(BIN_TO_UUID(product_id), '-', ''), name, price, quantity, description, image, REPLACE(BIN_TO_UUID(category_id), '-', ''), 
	created_at, updated_at FROM products WHERE id = ?`

	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)

	defer rows.Close()

	product := domain.Product{}

	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.Image, &product.CategoryId, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicIfError(err)

		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := `SELECT REPLACE(BIN_TO_UUID(product_id), '-', ''), name, price, quantity, description, image, REPLACE(BIN_TO_UUID(category_id), '-', ''), 
	created_at, updated_at FROM products`

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	products := []domain.Product{}

	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.Image, &product.CategoryId, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicIfError(err)

		products = append(products, product)
	}

	return products
}

func (repository *ProductRepositoryImpl) SaveMany(ctx context.Context, tx *sql.Tx, products []domain.Product) []domain.Product {
	SQL := "INSERT INTO products(product_id, name, price, quantity, description, image, category_id, created_at) values (UUID_TO_BIN(UUID()),?,?,?,?,?,UUID_TO_BIN(?),?)"

	stmt, err := tx.PrepareContext(ctx, SQL)
	helper.PanicIfError(err)

	defer stmt.Close()

	for _, product := range products {
		result, err := stmt.ExecContext(ctx, product.Name, product.Price, product.Quantity, product.Description, product.Image, product.CategoryId, product.CreatedAt)
		helper.PanicIfError(err)

		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		product.ProductId = int(id)
	}

	return products
}

func (repository *ProductRepositoryImpl) DeleteAll(ctx context.Context, tx *sql.Tx) {
	SQL := "TRUNCATE products"

	_, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
}
