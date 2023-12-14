package primary

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/projectContainsUser"
)

type IProjectContainsUserServices interface {
	Create(projectID uuid.UUID, projectContainsUserObject projectContainsUser.ProjectContainsUser) errors.Error
	List(projectID uuid.UUID) ([]projectContainsUser.ProjectContainsUser, errors.Error)
	InverseList(projectID uuid.UUID) ([]projectContainsUser.ProjectContainsUser, errors.Error)
	Delete(projectID, userID uuid.UUID) errors.Error
}
