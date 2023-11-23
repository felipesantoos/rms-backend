package primary

import (
	"rms-backend/src/core/domain/errors"
	_type "rms-backend/src/core/domain/type"
)

type ITypeServices interface {
	List() ([]_type.Type, errors.Error)
}
