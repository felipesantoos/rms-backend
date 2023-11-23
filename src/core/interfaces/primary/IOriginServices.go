package primary

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/origin"
)

type IOriginServices interface {
	List() ([]origin.Origin, errors.Error)
}
