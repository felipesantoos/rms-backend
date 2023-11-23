package services

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/type"
	"rms-backend/src/core/interfaces/primary"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
)

type typeServices struct {
	typeRepository repository.ITypeRepository
}

func NewTypeServices(typeRepository repository.ITypeRepository) primary.ITypeServices {
	return &typeServices{typeRepository: typeRepository}
}

func (this *typeServices) List() ([]_type.Type, errors.Error) {
	types, err := this.typeRepository.List()
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return types, nil
}
