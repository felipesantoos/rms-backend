package services

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/user"
	"rms-backend/src/core/interfaces/primary"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
)

type userServices struct {
	userRepository repository.IUserRepository
}

func NewUserServices(userRepository repository.IUserRepository) primary.IUserServices {
	return &userServices{userRepository: userRepository}
}

func (this *userServices) GetByEmail(email string) (user.User, errors.Error) {
	userObject, err := this.userRepository.GetByEmail(email)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return userObject, nil
}
