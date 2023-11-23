package response

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/requirement"
	"time"
)

type Requirement struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserStory   string    `json:"user_story"`
	TypeID      uuid.UUID `json:"type_id"`
	OriginID    uuid.UUID `json:"origin_id"`
	PriorityID  uuid.UUID `json:"priority_id"`
	ProjectID   uuid.UUID `json:"project_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type requirementBuilder struct{}

func RequirementBuilder() *requirementBuilder {
	return &requirementBuilder{}
}

func (*requirementBuilder) BuildFromDomain(data requirement.Requirement) Requirement {
	return Requirement{
		ID:          data.ID(),
		Title:       data.Title(),
		Description: data.Description(),
		UserStory:   data.UserStory(),
		TypeID:      data.TypeID(),
		OriginID:    data.OriginID(),
		PriorityID:  data.PriorityID(),
		ProjectID:   data.ProjectID(),
		CreatedAt:   data.CreatedAt(),
		UpdatedAt:   data.UpdatedAt(),
	}
}

func (builder *requirementBuilder) BuildFromDomainList(list []requirement.Requirement) []Requirement {
	result := []Requirement{}
	for _, item := range list {
		result = append(result, builder.BuildFromDomain(item))
	}
	return result
}
