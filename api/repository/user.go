package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/spark-tokyo/atlas/api/entity"
	"github.com/spark-tokyo/atlas/config"
	"github.com/spark-tokyo/atlas/ent"
)

var _ IFUserRepository = (*UserRepository)(nil)

type IFUserRepository interface {
	Get(ctx context.Context, tx *ent.Tx, id string) (*entity.User, error)
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

func (r *UserRepository) Get(ctx context.Context, tx *ent.Tx, id string) (*entity.User, error) {
	tx, err := FetchTx(tx)
	if err != nil {
		return nil, err
	}

	//! 一時的な実装 DB のID を String にする
	base := 10
	bitSize := 64
	number, err := strconv.ParseInt(id, base, bitSize)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted number:", number)
	}
	//!

	user, err := tx.User.Get(ctx, int(number))
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

func FetchTx(tx *ent.Tx) (*ent.Tx, error) {
	if tx == nil {
		return nil, errors.New("tx is nil")
	}
	return tx, nil
}
