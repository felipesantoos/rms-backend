package services

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/project"
	"rms-backend/src/core/interfaces/primary"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/services/filters"
)

type projectServices struct {
	projectRepository repository.IProjectRepository
}

func NewProjectServices(projectRepository repository.IProjectRepository) primary.IProjectServices {
	return &projectServices{projectRepository: projectRepository}
}

func (this *projectServices) Create(projectObject project.Project) (*uuid.UUID, errors.Error) {
	id, err := this.projectRepository.Create(projectObject)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return id, nil
}

func (this *projectServices) List(projectFilters filters.ProjectFilters) ([]project.Project, errors.Error) {
	projects, err := this.projectRepository.List(projectFilters)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return projects, nil
}

func (this *projectServices) Get(id uuid.UUID) (project.Project, errors.Error) {
	projectObject, err := this.projectRepository.Get(id)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return projectObject, nil
}

func (this *projectServices) Update(projectObject project.Project) errors.Error {
	err := this.projectRepository.Update(projectObject)
	if err != nil {
		return logger.LogCustomError(err)
	}
	return nil
}

func (this *projectServices) Delete(id uuid.UUID) errors.Error {
	err := this.projectRepository.Delete(id)
	if err != nil {
		return logger.LogCustomError(err)
	}
	return nil
}
