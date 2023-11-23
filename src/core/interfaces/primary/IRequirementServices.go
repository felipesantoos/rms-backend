package primary

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/requirement"
	"rms-backend/src/core/services/filters"
)

type IRequirementServices interface {
	Create(requirementObject requirement.Requirement) (*uuid.UUID, errors.Error)
	List(requirementFilters filters.RequirementFilters) ([]requirement.Requirement, errors.Error)
	Get(id uuid.UUID) (requirement.Requirement, errors.Error)
	Update(requirementObject requirement.Requirement) errors.Error
	Delete(id uuid.UUID) errors.Error
}
