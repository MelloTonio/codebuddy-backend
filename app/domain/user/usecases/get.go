package usecases

import (
	"github.com/mellotonio/desafiogo/app/domain/user"
)

func (du UserUsecase) Get(nickName string) (user.User, error) {
	dbUser, err := du.UserRepository.Get(nickName)
	if err != nil {
		return user.User{}, err
	}

	return dbUser, nil
}
