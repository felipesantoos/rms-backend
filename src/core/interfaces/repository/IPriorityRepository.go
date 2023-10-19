package repository

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/priority"
)

type IPriorityRepository interface {
	List() ([]priority.Priority, errors.Error)
}
