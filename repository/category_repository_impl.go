package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories(category_id, name, created_at) values (UUID_TO_BIN(UUID()),?,?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name, category.CreatedAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE categories SET name = ?, updated_at = ? WHERE REPLACE(BIN_TO_UUID(category_id), '-', '') = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.UpdatedAt, category.CategoryId)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM categories WHERE REPLACE(BIN_TO_UUID(category_id), '-', '') = ?"
	_, err := tx.ExecContext(ctx, SQL, category.CategoryId)
	helper.PanicIfError(err)

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId string) (domain.Category, error) {
	SQL := "SELECT REPLACE(BIN_TO_UUID(category_id), '-', ''), name,created_at, updated_at FROM categories WHERE REPLACE(BIN_TO_UUID(category_id), '-', '') = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.CategoryId, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT REPLACE(BIN_TO_UUID(category_id), '-', ''), name,created_at, updated_at FROM categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	categories := []domain.Category{}
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.CategoryId, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		helper.PanicIfError(err)

		categories = append(categories, category)

	}
	return categories
}

func (repository *CategoryRepositoryImpl) FindId(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT REPLACE(BIN_TO_UUID(category_id), '-', ''), name,created_at, updated_at FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.CategoryId, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) SaveMany(ctx context.Context, tx *sql.Tx, categories []domain.Category) []domain.Category {
	SQL := "INSERT INTO categories(category_id, name, created_at) values (UUID_TO_BIN(UUID()),?,?)"

	stmt, err := tx.PrepareContext(ctx, SQL)
	helper.PanicIfError(err)

	defer stmt.Close()

	for _, category := range categories {
		result, err := stmt.ExecContext(ctx, category.Name, category.CreatedAt)
		helper.PanicIfError(err)

		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		category.Id = int(id)
	}

	return categories
}

func (repository *CategoryRepositoryImpl) DeleteAll(ctx context.Context, tx *sql.Tx) {
	SQL := "DELETE FROM categories WHERE created_at != 1682954749732;"
	_, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
}
