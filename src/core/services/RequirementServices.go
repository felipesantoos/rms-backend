package services

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/requirement"
	"rms-backend/src/core/interfaces/primary"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/services/filters"
)

type requirementServices struct {
	requirementRepository repository.IRequirementRepository
}

func NewRequirementServices(requirementRepository repository.IRequirementRepository) primary.IRequirementServices {
	return &requirementServices{requirementRepository: requirementRepository}
}

func (this *requirementServices) Create(requirementObject requirement.Requirement) (*uuid.UUID, errors.Error) {
	lastCode, err := this.requirementRepository.GetLastCodeByProject(requirementObject.ProjectID())
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	requirementObject.SetCode(*lastCode + 1)
	id, err := this.requirementRepository.Create(requirementObject)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return id, nil
}

func (this *requirementServices) List(requirementFilters filters.RequirementFilters) ([]requirement.Requirement, errors.Error) {
	requirements, err := this.requirementRepository.List(requirementFilters)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return requirements, nil
}

func (this *requirementServices) Get(id uuid.UUID) (requirement.Requirement, errors.Error) {
	requirementObject, err := this.requirementRepository.Get(id)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return requirementObject, nil
}

func (this *requirementServices) Update(requirementObject requirement.Requirement) errors.Error {
	err := this.requirementRepository.Update(requirementObject)
	if err != nil {
		return logger.LogCustomError(err)
	}
	return nil
}

func (this *requirementServices) Delete(id uuid.UUID) errors.Error {
	err := this.requirementRepository.Delete(id)
	if err != nil {
		return logger.LogCustomError(err)
	}
	return nil
}
