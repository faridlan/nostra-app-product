package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	panic("not implemented") // TODO: Implement
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	panic("not implemented") // TODO: Implement
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	panic("not implemented") // TODO: Implement
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId string) (domain.Category, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	panic("not implemented") // TODO: Implement
}

func (repository *CategoryRepositoryImpl) FindId(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	panic("not implemented") // TODO: Implement
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

		category.CateogoryId = int(id)
	}

	return categories
}

func (repository *CategoryRepositoryImpl) DeleteAll(ctx context.Context, tx *sql.Tx) {
	SQL := "DELETE FROM categories WHERE created_at != 1682954749732;"
	_, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
}
