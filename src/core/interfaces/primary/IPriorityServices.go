package primary

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/priority"
)

type IPriorityServices interface {
	List() ([]priority.Priority, errors.Error)
}
