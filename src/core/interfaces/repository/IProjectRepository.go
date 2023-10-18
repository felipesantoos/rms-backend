package repository

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/project"
)

type IProjectRepository interface {
	Create(projectObject project.Project) (*uuid.UUID, errors.Error)
}
