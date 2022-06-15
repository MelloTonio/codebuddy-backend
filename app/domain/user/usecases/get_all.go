package usecases

import (
	"github.com/mellotonio/desafiogo/app/domain/user"
)

type UserUsecase struct {
	UserRepository user.Repository
}

func NewUserUsecase(userRepository user.Repository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}

func (du UserUsecase) GetAll() ([]user.User, error) {
	users, err := du.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}
