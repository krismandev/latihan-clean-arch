package repository

import (
	"agit-test/helper"
	"agit-test/model/domain"
	"context"
	"database/sql"
	"time"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {

	var baseIn string
	var baseParam []interface{}

	helper.AppendComma(&baseIn, &baseParam, "?", user.Username)
	helper.AppendComma(&baseIn, &baseParam, "?", user.Password)
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	helper.AppendComma(&baseIn, &baseParam, "?", timeNow)

	query := "INSERT INTO users (username, password, created_at) VALUES (" + baseIn + ")"

	res, err := tx.ExecContext(ctx, query, baseParam...)
	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	user.CreatedAt = timeNow

	return user
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) domain.User {
	var user domain.User

	query := "SELECT id, username, password, created_at, updated_at, deleted_at FROM users WHERE username = ?"
	rows, err := tx.QueryContext(ctx, query, username)
	defer rows.Close()
	helper.PanicIfError(err)

	if rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		// helper.PanicIfError(err)
		// return user, nil
	}

	return user

}
