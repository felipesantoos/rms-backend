package adapters

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/knowledgeArea"
)

type ResourcesAdapter interface {
	ListKnowledgeAreas() ([]knowledgeArea.KnowledgeArea, errors.Error)
}
