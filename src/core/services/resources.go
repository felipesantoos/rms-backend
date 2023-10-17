package services

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/knowledgeArea"
	"rms-backend/src/core/interfaces/adapters"
	"rms-backend/src/core/interfaces/usecases"
)

type resourcesService struct {
	services adapters.ResourcesAdapter
}

func NewResourcesService(adapter adapters.ResourcesAdapter) usecases.ResourcesUseCase {
	return &resourcesService{adapter}
}

func (this *resourcesService) ListKnowledgeAreas() ([]knowledgeArea.KnowledgeArea, errors.Error) {
	return this.services.ListKnowledgeAreas()
}
