package primary

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/project"
)

type IProjectServices interface {
	Create(projectObject project.Project) (*uuid.UUID, errors.Error)
}
