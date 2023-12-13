package repository

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/user"
)

type IUserRepository interface {
	GetByEmail(email string) (user.User, errors.Error)
}
