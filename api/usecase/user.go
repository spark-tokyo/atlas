package usecase

import (
	"context"

	"github.com/spark-tokyo/atlas/api/entity"
	"github.com/spark-tokyo/atlas/api/infra"
	"github.com/spark-tokyo/atlas/api/repository"
	"github.com/spark-tokyo/atlas/ent"
	"github.com/spark-tokyo/atlas/tx"
)

var _ IFUserUsecase = (*UserUsecase)(nil)

type IFUserUsecase interface {
	Get(ctx context.Context, id string) (*User, error)
}

type UserUsecase struct {
	userRepository repository.IFUserRepository
	txManager      tx.IFTxManager
	ent            *infra.Ent
}

func NewUserUsecase(
	userRepository *repository.UserRepository,
	tx *tx.TxManager,
	ent *infra.Ent,
) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
		txManager:      tx,
		ent:            ent,
	}
}

type User struct {
	Id    string
	Name  string
	Email string
}

func (u *UserUsecase) Get(ctx context.Context, id string) (*User, error) {
	var userEntity *entity.User
	if err := u.txManager.WitTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		var err error
		userEntity, err = u.userRepository.Get(ctx, tx, id)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return toUserUsecase(userEntity), nil
}

func toUserUsecase(input *entity.User) *User {
	return &User{
		Id:    input.Id,
		Name:  input.Name,
		Email: input.Email,
	}
}
