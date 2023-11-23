package requirement

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/messages"
	"strings"
	"time"
)

type builder struct {
	fields        []string
	errorMessages []string
	requirement   *requirement
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		requirement:   &requirement{},
	}
}

func (this *builder) WithID(id uuid.UUID) *builder {
	if id.ID() == 0 {
		this.fields = append(this.fields, messages.RequirementID)
		this.errorMessages = append(this.errorMessages, messages.RequirementIDCannotBeEmpty)
		return this
	}
	this.requirement.id = id
	return this
}

func (this *builder) WithCode(code int) *builder {
	if code == 0 {
		this.fields = append(this.fields, messages.RequirementCode)
		this.errorMessages = append(this.errorMessages, messages.RequirementCodeCannotBeEmpty)
		return this
	}
	this.requirement.code = code
	return this
}

func (this *builder) WithTitle(title string) *builder {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		this.fields = append(this.fields, messages.RequirementTitle)
		this.errorMessages = append(this.errorMessages, messages.RequirementTitleCannotBeEmpty)
		return this
	}
	this.requirement.title = title
	return this
}

func (this *builder) WithDescription(description string) *builder {
	this.requirement.description = description
	return this
}

func (this *builder) WithUserStory(userStory string) *builder {
	this.requirement.userStory = userStory
	return this
}

func (this *builder) WithTypeID(typeID uuid.UUID) *builder {
	if typeID.ID() == 0 {
		this.fields = append(this.fields, messages.RequirementTypeID)
		this.errorMessages = append(this.errorMessages, messages.RequirementTypeIDCannotBeEmpty)
		return this
	}
	this.requirement.typeID = typeID
	return this
}

func (this *builder) WithOriginID(originID uuid.UUID) *builder {
	if originID.ID() == 0 {
		this.fields = append(this.fields, messages.RequirementOriginID)
		this.errorMessages = append(this.errorMessages, messages.RequirementOriginIDCannotBeEmpty)
		return this
	}
	this.requirement.originID = originID
	return this
}

func (this *builder) WithPriorityID(priorityID uuid.UUID) *builder {
	if priorityID.ID() == 0 {
		this.fields = append(this.fields, messages.RequirementPriorityID)
		this.errorMessages = append(this.errorMessages, messages.RequirementPriorityIDCannotBeEmpty)
		return this
	}
	this.requirement.priorityID = priorityID
	return this
}

func (this *builder) WithProjectID(projectID uuid.UUID) *builder {
	if projectID.ID() == 0 {
		this.fields = append(this.fields, messages.RequirementProjectID)
		this.errorMessages = append(this.errorMessages, messages.RequirementProjectIDCannotBeEmpty)
		return this
	}
	this.requirement.projectID = projectID
	return this
}

func (this *builder) WithCreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		this.fields = append(this.fields, messages.RequirementCreatedAt)
		this.errorMessages = append(this.errorMessages, messages.RequirementCreatedAtCannotBeEmpty)
		return this
	}
	this.requirement.createdAt = createdAt
	return this
}

func (this *builder) WithUpdatedAt(updatedAt time.Time) *builder {
	this.requirement.updatedAt = updatedAt
	return this
}

func (this *builder) WithDeletedAt(deletedAt *time.Time) *builder {
	this.requirement.deletedAt = deletedAt
	return this
}

func (this *builder) Build() (*requirement, errors.Error) {
	if len(this.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(this.errorMessages, map[string]interface{}{
			messages.Fields: this.fields})
	}
	return this.requirement, nil
}
