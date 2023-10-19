package services

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/project"
	"rms-backend/src/core/interfaces/primary"
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

func (this *projectServices) List() ([]project.Project, errors.Error) {
	return this.projectRepository.List()
}

func (this *projectServices) Get(id uuid.UUID) (project.Project, errors.Error) {
	return this.projectRepository.Get(id)
}

func (this *projectServices) Update(projectObject project.Project) errors.Error {
	return this.projectRepository.Update(projectObject)
}

func (this *projectServices) Delete(id uuid.UUID) errors.Error {
	return nil
}
