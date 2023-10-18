package services

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/project"
	primary "rms-backend/src/core/interfaces/primary"
	"rms-backend/src/core/interfaces/repository"
)

type projectServices struct {
	projectRepository repository.IProjectRepository
}

func NewProjectServices(projectRepository repository.IProjectRepository) primary.IProjectServices {
	return &projectServices{projectRepository: projectRepository}
}

func (this *projectServices) Create(projectObject project.Project) (*uuid.UUID, errors.Error) {
	return this.projectRepository.Create(projectObject)
}
