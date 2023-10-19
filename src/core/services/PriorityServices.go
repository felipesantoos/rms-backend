package services

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/priority"
	"rms-backend/src/core/interfaces/primary"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
)

type priorityServices struct {
	priorityRepository repository.IPriorityRepository
}

func NewPriorityServices(priorityRepository repository.IPriorityRepository) primary.IPriorityServices {
	return &priorityServices{priorityRepository: priorityRepository}
}

func (this *priorityServices) List() ([]priority.Priority, errors.Error) {
	priorities, err := this.priorityRepository.List()
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return priorities, nil
}
