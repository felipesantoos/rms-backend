package primary

import (
	"rms-backend/src/core/domain/authorization"
	"rms-backend/src/core/domain/credentials"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/user"
)

type IAuthServices interface {
	Login(credentials credentials.Credentials) (authorization.Authorization, errors.Error)
	SignUp(_user user.User) (authorization.Authorization, errors.Error)
}
