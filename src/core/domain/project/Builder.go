package project

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
	project       *project
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		project:       &project{},
	}
}

func (this *builder) WithID(id uuid.UUID) *builder {
	if id.ID() == 0 {
		this.fields = append(this.fields, messages.ProjectID)
		this.errorMessages = append(this.errorMessages, messages.ProjectIDCannotBeEmpty)
		return this
	}
	this.project.id = id
	return this
}

func (this *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		this.fields = append(this.fields, messages.ProjectName)
		this.errorMessages = append(this.errorMessages, messages.ProjectNameCannotBeEmpty)
		return this
	}
	this.project.name = name
	return this
}

func (this *builder) WithAlias(alias string) *builder {
	alias = strings.TrimSpace(alias)
	if len(alias) == 0 {
		this.fields = append(this.fields, messages.ProjectAlias)
		this.errorMessages = append(this.errorMessages, messages.ProjectAliasCannotBeEmpty)
		return this
	}
	this.project.alias = alias
	return this
}

func (this *builder) WithDescription(description string) *builder {
	this.project.description = description
	return this
}

func (this *builder) WithIsActive(isActive bool) *builder {
	this.project.isActive = isActive
	return this
}

func (this *builder) WithCreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		this.fields = append(this.fields, messages.ProjectCreatedAt)
		this.errorMessages = append(this.errorMessages, messages.ProjectCreatedAtCannotBeEmpty)
		return this
	}
	this.project.createdAt = createdAt
	return this
}

func (this *builder) WithUpdatedAt(updatedAt time.Time) *builder {
	this.project.updatedAt = updatedAt
	return this
}

func (this *builder) WithDeletedAt(deletedAt time.Time) *builder {
	this.project.deletedAt = deletedAt
	return this
}

func (this *builder) Build() (*project, errors.Error) {
	if len(this.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(this.errorMessages, map[string]interface{}{
			messages.Fields: this.fields})
	}
	return this.project, nil
}
