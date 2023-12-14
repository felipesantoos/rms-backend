package services

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/projectContainsUser"
	"rms-backend/src/core/interfaces/primary"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
)

type projectContainsUserServices struct {
	projectContainsUserRepository repository.IProjectContainsUserRepository
}

func NewProjectContainsUserServices(projectContainsUserRepository repository.IProjectContainsUserRepository) primary.IProjectContainsUserServices {
	return &projectContainsUserServices{projectContainsUserRepository: projectContainsUserRepository}
}

func (this *projectContainsUserServices) Create(projectID uuid.UUID, projectContainsUserObject projectContainsUser.ProjectContainsUser) errors.Error {
	isCollaboratorAlreadyRegistered, err := this.projectContainsUserRepository.IsCollaborator(projectID, projectContainsUserObject.UserID())
	if err != nil {
		return logger.LogCustomError(err)
	}

	if isCollaboratorAlreadyRegistered {
		err = this.projectContainsUserRepository.ReactivateCollaborator(projectID, projectContainsUserObject.UserID())
		if err != nil {
			return logger.LogCustomError(err)
		}
		return nil
	}

	err = this.projectContainsUserRepository.Create(projectID, projectContainsUserObject)
	if err != nil {
		return logger.LogCustomError(err)
	}
	return nil
}

func (this *projectContainsUserServices) List(projectID uuid.UUID) ([]projectContainsUser.ProjectContainsUser, errors.Error) {
	projectContainsUsers, err := this.projectContainsUserRepository.List(projectID)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return projectContainsUsers, nil
}

func (this *projectContainsUserServices) InverseList(projectID uuid.UUID) ([]projectContainsUser.ProjectContainsUser, errors.Error) {
	projectContainsUsers, err := this.projectContainsUserRepository.InverseList(projectID)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return projectContainsUsers, nil
}

func (this *projectContainsUserServices) Delete(projectID, userID uuid.UUID) errors.Error {
	err := this.projectContainsUserRepository.Delete(projectID, userID)
	if err != nil {
		return logger.LogCustomError(err)
	}
	return nil
}
