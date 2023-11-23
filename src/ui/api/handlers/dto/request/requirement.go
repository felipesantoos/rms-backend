package request

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/requirement"
	"rms-backend/src/core/logger"
)

type Requirement struct {
	Title       string    `json:"title" example:"Requirement A"`
	Description string    `json:"description" example:"Description of the Requirement A."`
	UserStory   string    `json:"user_story" example:"User story of the Requirement A."`
	TypeID      uuid.UUID `json:"type_id" example:"9e57a3e7-8be3-4c0e-8a0d-6ae7cde4b87c"`
	OriginID    uuid.UUID `json:"origin_id" example:"c1d8e2a4-88ea-43e6-bb90-5711b34b7a8f"`
	PriorityID  uuid.UUID `json:"priority_id" example:"f3ab3bf0-76f1-4ad7-8c3b-1f93eac9d4a0"`
	ProjectID   uuid.UUID `json:"project_id" example:"38b27910-e17f-4b8d-bc3f-85ca1c76e3db"`
}

func (this *Requirement) ToDomain() (requirement.Requirement, errors.Error) {
	requirementObject, validationError := requirement.NewBuilder().WithTitle(this.Title).
		WithDescription(this.Description).WithUserStory(this.UserStory).WithTypeID(this.TypeID).
		WithOriginID(this.OriginID).WithPriorityID(this.PriorityID).WithProjectID(this.ProjectID).Build()
	if validationError != nil {
		return nil, logger.LogCustomError(validationError)
	}
	return requirementObject, nil
}
