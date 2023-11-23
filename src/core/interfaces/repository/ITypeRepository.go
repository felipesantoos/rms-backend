package repository

import (
	"rms-backend/src/core/domain/errors"
	_type "rms-backend/src/core/domain/type"
)

type ITypeRepository interface {
	List() ([]_type.Type, errors.Error)
}
