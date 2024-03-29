package services

import (
	"rms-backend/src/core/domain/authorization"
	"rms-backend/src/core/domain/credentials"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/user"
	"rms-backend/src/core/interfaces/primary"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/messages"
	"rms-backend/src/core/utils/encrypt"
)

type authServices struct {
	authRepository repository.IAuthRepository
}

func NewAuthServices(authRepository repository.IAuthRepository) primary.IAuthServices {
	return &authServices{authRepository: authRepository}
}

func (instance authServices) Login(_credentials credentials.Credentials) (authorization.Authorization, errors.Error) {
	_user, err := instance.authRepository.Login(_credentials)
	if err != nil {
		logger.LogCustomError(err)
		return nil, err
	}

	passwordsMatch := encrypt.PasswordsMatch(_credentials.Password(), _user.Salt(), _user.Password())
	if !passwordsMatch {
		return nil, errors.NewValidationFromString(messages.TheProvidedCredentialsAreInvalid)
	}

	_authorization, authorizationError := authorization.NewFromAccount(_user)
	if authorizationError != nil {
		logger.LogCustomError(authorizationError)
		return nil, authorizationError
	}

	return _authorization, nil
}

func (instance authServices) SignUp(_user user.User) (authorization.Authorization, errors.Error) {
	password, salt, encryptError := encrypt.EncryptPassword(_user.Password())
	if encryptError != nil {
		logger.LogNativeError(encryptError)
		return nil, errors.NewInternal(encryptError)
	}
	_user.SetPassword(password)
	_user.SetSalt(salt)
	id, err := instance.authRepository.SignUp(_user)
	if err != nil {
		logger.LogCustomError(err)
		return nil, err
	}
	_user.SetID(*id)
	_authorization, authorizationError := authorization.NewFromAccount(_user)
	if authorizationError != nil {
		logger.LogCustomError(authorizationError)
		return nil, authorizationError
	}

	return _authorization, nil
}
