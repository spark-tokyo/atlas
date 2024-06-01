package repository

import (
	"context"
	"fmt"

	"github.com/spark-tokyo/atlas/api/entity"
	"github.com/spark-tokyo/atlas/config"
	"github.com/spark-tokyo/atlas/ent"
)

var _ IFUserRepository = (*UserRepository)(nil)

type IFUserRepository interface {
	Get(ctx context.Context, tx *ent.Tx) (*entity.User, error)
}

type UserRepository struct {
	config *config.Config
}

func NewUserRepository(
	cfg *config.Config,
) *UserRepository {
	return &UserRepository{
		config: cfg,
	}
}

func (r *UserRepository) Get(ctx context.Context, tx *ent.Tx) (*entity.User, error) {
	user, err := tx.User.Get(ctx, 1)
	if err != nil {
		return nil, err
	}
	res := &entity.User{
		Id:    fmt.Sprint(user.ID),
		Name:  "Name",
		Email: "Email",
	}
	return res, nil
}
