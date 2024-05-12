package repository

import (
	"context"

	"atlas/api/entity"
)

type IFUserRepository interface {
	Get(ctx context.Context) (*entity.User, error)
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) Get(ctx context.Context) (*entity.User, error) {
	res := &entity.User{
		Id:    "1",
		Name:  "Name",
		Email: "Email",
	}
	return res, nil
}
