package primary

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/project"
)

type IProjectServices interface {
	Create(projectObject project.Project) (*uuid.UUID, errors.Error)
	List() ([]project.Project, errors.Error)
	Get(id uuid.UUID) (project.Project, errors.Error)
	Update(projectObject project.Project) errors.Error
	Delete(id uuid.UUID) errors.Error
}
