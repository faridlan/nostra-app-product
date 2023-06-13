package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/nostra-api-product/helper"
	"github.com/faridlan/nostra-api-product/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(user_id, username, password, email, image, role_id, created_at) values (UUID_TO_BIN(UUID()),?,?,?,?,UUID_TO_BIN(?),?)"

	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.Email, user.Image, user.Role.RoleId, user.CreatedAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET username = ?, email = ?, image = ?, role_id = UUID_TO_BIN(?), updated_at = ? WHERE REPLACE(BIN_TO_UUID(user_id), '-', '') = ?"

	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Image, user.Role.RoleId, user.UpdatedAt, user.UserId)
	helper.PanicIfError(err)

	return user

}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId string) (domain.User, error) {
	SQL := `SELECT REPLACE(BIN_TO_UUID(u.user_id), '-', '') as user_id, u.username, u.email, u.image, REPLACE(BIN_TO_UUID(r.role_id), '-', '') as role_id, r.name, r.created_at, r.updated_at, u.created_at, u.updated_at 
	FROM users AS u 
	INNER JOIN roles AS r ON (r.role_id = u.role_id)
	WHERE REPLACE(BIN_TO_UUID(user_id), '-', '') = ?`

	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	user := domain.User{}

	if rows.Next() {
		err := rows.Scan(&user.UserId, &user.Username, &user.Email, &user.Image, &user.Role.RoleId, &user.Role.Name, &user.Role.CreatedAt, &user.Role.UpdatedAt, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindId(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := `SELECT REPLACE(BIN_TO_UUID(u.user_id), '-', '') as user_id, u.username, u.email, u.image, REPLACE(BIN_TO_UUID(r.role_id), '-', '') as role_id, r.name ,r.created_at, r.updated_at, u.created_at, u.updated_at 
	FROM users AS u 
	INNER JOIN roles AS r ON (r.role_id = u.role_id)
	WHERE u.id = ?`

	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	user := domain.User{}

	if rows.Next() {
		err := rows.Scan(&user.UserId, &user.Username, &user.Email, &user.Image, &user.Role.RoleId, &user.Role.Name, &user.Role.CreatedAt, &user.Role.UpdatedAt, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)

		return user, nil
	} else {
		return user, errors.New("user not found")
	}

}

func (repository *UserRepositoryImpl) FindName(ctx context.Context, tx *sql.Tx, username string) (string, error) {
	SQL := `SELECT username FROM users WHERE username = ?`

	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)

	defer rows.Close()

	user := domain.User{}

	if !rows.Next() {
		return user.Username, nil
	} else {
		return user.Username, errors.New("username has been taken")
	}

}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := `SELECT REPLACE(BIN_TO_UUID(u.user_id), '-', '') as user_id, u.username, u.email, u.image, REPLACE(BIN_TO_UUID(r.role_id), '-', '') as role_id, r.name, r.created_at, r.updated_at, u.created_at, u.updated_at 
	FROM users AS u 
	INNER JOIN roles AS r ON (r.role_id = u.role_id)
	ORDER BY u.created_at DESC`

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	users := []domain.User{}

	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.UserId, &user.Username, &user.Email, &user.Image, &user.Role.RoleId, &user.Role.Name, &user.Role.CreatedAt, &user.Role.UpdatedAt, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users

}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := `SELECT REPLACE(BIN_TO_UUID(u.user_id), '-', '') as user_id, u.username, u.password, u.email, u.image, REPLACE(BIN_TO_UUID(r.role_id), '-', '') as role_id, r.name, r.created_at, r.updated_at, u.created_at, u.updated_at
	FROM users AS u 
	INNER JOIN roles AS r ON (r.role_id = u.role_id)
	WHERE u.username = ?`

	rows, err := tx.QueryContext(ctx, SQL, user.Username)
	helper.PanicIfError(err)

	defer rows.Close()

	userModel := domain.User{}

	if rows.Next() {
		err := rows.Scan(&userModel.UserId, &userModel.Username, &userModel.Password, &userModel.Email, &userModel.Image, &userModel.Role.RoleId, &userModel.Role.Name, &userModel.Role.CreatedAt, &userModel.Role.UpdatedAt, &userModel.CreatedAt, &userModel.UpdatedAt)
		helper.PanicIfError(err)

		return userModel, nil
	} else {
		return userModel, errors.New("username or password incorect")
	}
}

func (repository *UserRepositoryImpl) SaveMany(ctx context.Context, tx *sql.Tx, users []domain.User) []domain.User {
	SQL := "INSERT INTO users(user_id, username, password, email, image, role_id, created_at) values (UUID_TO_BIN(UUID()),?,?,?,?,UUID_TO_BIN(?),?)"

	stmt, err := tx.PrepareContext(ctx, SQL)
	helper.PanicIfError(err)

	defer stmt.Close()

	for _, user := range users {
		result, err := stmt.ExecContext(ctx, user.Username, user.Password, user.Email, user.Image, user.Role.RoleId, user.CreatedAt)
		helper.PanicIfError(err)

		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		user.Id = int(id)
	}

	return users
}

func (repository *UserRepositoryImpl) DeleteAll(ctx context.Context, tx *sql.Tx) {
	SQL := "TRUNCATE users"

	_, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
}
