package repository

import (
	"agit-test/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) domain.User
}
