package repository

import (
	"rms-backend/src/core/domain/credentials"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/user"
)

type IAuthRepository interface {
	Login(credentials credentials.Credentials) (user.User, errors.Error)
}
