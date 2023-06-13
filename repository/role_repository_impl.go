package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/domain"
)

type RoleRepositoryImpl struct {
}

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

func (repository *RoleRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	SQL := "INSERT INTO roles(role_id, name, created_at) values (UUID_TO_BIN(UUID()),?,?)"
	result, err := tx.ExecContext(ctx, SQL, role.Name, role.CreatedAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	role.Id = int(id)

	return role
}

func (repository *RoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	SQL := "UPDATE roles SET name = ?, updated_at = ? WHERE REPLACE(BIN_TO_UUID(role_id), '-', '') = ?"
	_, err := tx.ExecContext(ctx, SQL, role.Name, role.UpdatedAt, role.RoleId)
	helper.PanicIfError(err)

	return role
}

func (repository *RoleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, roleId string) (domain.Role, error) {
	SQL := "SELECT REPLACE(BIN_TO_UUID(role_id), '-', '') as role_id, name, created_at, updated_at FROM roles WHERE REPLACE(BIN_TO_UUID(role_id), '-', '') = ?"
	rows, err := tx.QueryContext(ctx, SQL, roleId)
	helper.PanicIfError(err)

	defer rows.Close()

	role := domain.Role{}

	if rows.Next() {
		err := rows.Scan(&role.RoleId, &role.Name, &role.CreatedAt, &role.UpdatedAt)
		helper.PanicIfError(err)

		return role, nil
	} else {
		return role, errors.New("role not found")
	}
}

func (repository *RoleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Role {
	SQL := "SELECT REPLACE(BIN_TO_UUID(role_id), '-', '') as role_id, name, created_at, updated_at FROM roles ORDER BY created_at DESC"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	roles := []domain.Role{}
	for rows.Next() {
		role := domain.Role{}
		err := rows.Scan(&role.RoleId, &role.Name, &role.CreatedAt, &role.UpdatedAt)
		helper.PanicIfError(err)

		roles = append(roles, role)
	}

	return roles
}

func (repository *RoleRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, name string) (domain.Role, error) {
	SQL := "SELECT REPLACE(BIN_TO_UUID(role_id), '-', '') as role_id, name, created_at, updated_at FROM roles WHERE name = ?"
	rows, err := tx.QueryContext(ctx, SQL, name)
	helper.PanicIfError(err)

	defer rows.Close()

	role := domain.Role{}

	if rows.Next() {
		err := rows.Scan(&role.RoleId, &role.Name, &role.CreatedAt, &role.UpdatedAt)
		helper.PanicIfError(err)

		return role, nil
	} else {
		return role, errors.New("role not found")
	}
}

func (repository *RoleRepositoryImpl) SaveMany(ctx context.Context, tx *sql.Tx, roles []domain.Role) []domain.Role {
	SQL := "INSERT INTO roles(role_id, name, created_at) values (UUID_TO_BIN(UUID()),?,?)"

	stmt, err := tx.PrepareContext(ctx, SQL)
	helper.PanicIfError(err)

	defer stmt.Close()

	for _, role := range roles {
		result, err := stmt.ExecContext(ctx, role.Name, role.CreatedAt)
		helper.PanicIfError(err)

		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		role.Id = int(id)
	}

	return roles
}

func (repository *RoleRepositoryImpl) DeleteAll(ctx context.Context, tx *sql.Tx) {
	// SQL := "DELETE FROM roles WHERE created_at != 1683382190182"
	// SQL := "TRUNCATE roles"
	SQL := "DELETE FROM roles"

	_, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
}
