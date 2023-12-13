package primary

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/user"
)

type IUserServices interface {
	GetByEmail(email string) (user.User, errors.Error)
}
