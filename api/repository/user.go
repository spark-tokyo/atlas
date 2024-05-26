package repository

import (
	"context"
	"fmt"

	"atlas/api/entity"
	"atlas/api/infra"
	"atlas/config"
)

var _ IFUserRepository = (*UserRepository)(nil)

type IFUserRepository interface {
	Get(ctx context.Context) (*entity.User, error)
}

type UserRepository struct {
	config *config.Config
	db     *infra.Ent
}

func NewUserRepository(
	cfg *config.Config,
	db *infra.Ent,
) *UserRepository {
	return &UserRepository{
		config: cfg,
		db:     db,
	}
}

func (r *UserRepository) Get(ctx context.Context) (*entity.User, error) {
	user, err := r.db.User.Get(ctx, 1)
	if err != nil {
		return nil, err
	}
	fmt.Print(r.config.Stage)
	res := &entity.User{
		Id:    fmt.Sprint(user.ID),
		Name:  "Name",
		Email: "Email",
	}
	return res, nil
}
