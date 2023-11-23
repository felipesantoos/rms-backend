package services

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/origin"
	"rms-backend/src/core/interfaces/primary"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
)

type originServices struct {
	originRepository repository.IOriginRepository
}

func NewOriginServices(originRepository repository.IOriginRepository) primary.IOriginServices {
	return &originServices{originRepository: originRepository}
}

func (this *originServices) List() ([]origin.Origin, errors.Error) {
	origins, err := this.originRepository.List()
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return origins, nil
}
