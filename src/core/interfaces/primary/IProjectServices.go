package primary

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/project"
	"rms-backend/src/core/services/filters"
)

type IProjectServices interface {
	Create(projectObject project.Project) (*uuid.UUID, errors.Error)
	List(projectFilters filters.ProjectFilters) ([]project.Project, errors.Error)
	Get(id uuid.UUID) (project.Project, errors.Error)
	Update(projectObject project.Project) errors.Error
	Delete(id uuid.UUID) errors.Error
}
