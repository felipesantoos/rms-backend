package repository

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/projectContainsUser"
)

type IProjectContainsUserRepository interface {
	Create(projectID uuid.UUID, projectContainsUserObject projectContainsUser.ProjectContainsUser) errors.Error
	List(projectID uuid.UUID) ([]projectContainsUser.ProjectContainsUser, errors.Error)
	InverseList(projectID uuid.UUID) ([]projectContainsUser.ProjectContainsUser, errors.Error)
	Delete(projectID, userID uuid.UUID) errors.Error

	IsCollaborator(projectID, userID uuid.UUID) (bool, errors.Error)
	ReactivateCollaborator(projectID, userID uuid.UUID) errors.Error
}
