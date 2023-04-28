package repository

import (
	"context"
	"database/sql"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/domain"
)

type UserRepositoryImpl struct {
}

func (repository *UserRepositoryImpl) SaveMany(ctx context.Context, tx *sql.Tx, users []domain.User) []domain.User {
	SQL := "INSERT INTO users(user_id, username, password, email, image, role_id, created_at) values (?,?,?,?,?,?,?)"

	stmt, err := tx.PrepareContext(ctx, SQL)
	helper.PanicIfError(err)

	defer stmt.Close()

	for _, user := range users {
		result, err := stmt.ExecContext(ctx, user.Id, user.Username, user.Password, user.Email, user.Image, user.RoleId, user.CreatedAt)
		helper.PanicIfError(err)

		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		user.UserId = int(id)
	}

	return users
}

func (repository *UserRepositoryImpl) DeleteAll(ctx context.Context, tx *sql.Tx) {
	SQL := "TRUNCATE users"

	_, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
}
