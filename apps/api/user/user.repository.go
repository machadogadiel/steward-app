package user

import "github.com/uptrace/bun"

type UserRepository struct {
	db *bun.DB
}

type repository interface {
	GetExistingUser(username string)
}

func NewRepo(db *bun.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
