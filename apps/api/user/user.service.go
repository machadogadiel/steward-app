package user

import (
	"context"
	"database/sql"
)

type UserService interface {
	GetExistingUser(id int, ctx context.Context) (User, error)
	CreateUser(user User, ctx context.Context) (sql.Result, error)
}

type service struct {
	repository UserRepository
}

func NewService(r UserRepository) UserService {
	return &service{
		repository: r,
	}
}

func (s *service) CreateUser(user User, ctx context.Context) (sql.Result, error) {
	return s.repository.CreateUser(user, ctx)
}

func (s *service) GetExistingUser(id int, ctx context.Context) (User, error) {
	return s.repository.GetExistingUser(id, ctx)
}
