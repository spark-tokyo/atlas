package repository

import (
	"context"

	"github.com/spark-tokyo/atlas/api/entity"
	"github.com/spark-tokyo/atlas/config"
	"github.com/spark-tokyo/atlas/ent"
)

var _ IFUserRepository = (*UserRepository)(nil)

type IFUserRepository interface {
	Get(
		ctx context.Context,
		tx *ent.Tx,
		id string,
	) (*entity.User, error)
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

func (r *UserRepository) Get(
	ctx context.Context,
	tx *ent.Tx,
	id string,
) (*entity.User, error) {
	tx, err := FetchTx(tx)
	if err != nil {
		return nil, err
	}

	user, err := tx.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return toUser(user), nil
}

func toUser(user *ent.User) *entity.User {
	return &entity.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
