package usecase

import (
	"context"

	"atlas/api/entity"
	"atlas/api/repository"
)

type IFUserUsecase interface {
	Get(ctx context.Context) (*User, error)
}

type UserUsecase struct {
	userRepository repository.IFUserRepository
}

func NewUserUsecase(
	userRepository *repository.IFUserRepository,
) *UserUsecase {
	return &UserUsecase{
		userRepository: *userRepository,
	}
}

type User struct {
	Id    string
	Name  string
	Email string
}

func (u *UserUsecase) Get(ctx context.Context) (*User, error) {
	userEntity, err := u.userRepository.Get(ctx)
	if err != nil {
		return nil, err
	}
	userUsecase := toUserUsecase(userEntity)
	return userUsecase, nil
}

func toUserUsecase(input *entity.User) *User {
	return &User{
		Id:    input.Id,
		Name:  input.Name,
		Email: input.Email,
	}
}
