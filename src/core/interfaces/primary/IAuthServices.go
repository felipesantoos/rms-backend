package primary

import (
	"rms-backend/src/core/domain/authorization"
	"rms-backend/src/core/domain/credentials"
	"rms-backend/src/core/domain/errors"
)

type IAuthServices interface {
	Login(credentials credentials.Credentials) (authorization.Authorization, errors.Error)
}
