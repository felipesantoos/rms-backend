package usecases

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/knowledgeArea"
)

type ResourcesUseCase interface {
	ListKnowledgeAreas() ([]knowledgeArea.KnowledgeArea, errors.Error)
}
