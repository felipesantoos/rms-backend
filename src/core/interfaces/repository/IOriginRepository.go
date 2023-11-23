package repository

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/origin"
)

type IOriginRepository interface {
	List() ([]origin.Origin, errors.Error)
}
