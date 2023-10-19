package priority

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
	priority      *priority
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		priority:      &priority{},
	}
}

func (this *builder) WithID(id uuid.UUID) *builder {
	if id.ID() == 0 {
		this.fields = append(this.fields, messages.PriorityID)
		this.errorMessages = append(this.errorMessages, messages.PriorityIDCannotBeEmpty)
		return this
	}
	this.priority.id = id
	return this
}

func (this *builder) WithLevel(level string) *builder {
	level = strings.TrimSpace(level)
	if len(level) == 0 {
		this.fields = append(this.fields, messages.PriorityLevel)
		this.errorMessages = append(this.errorMessages, messages.PriorityLevelCannotBeEmpty)
		return this
	}
	this.priority.level = level
	return this
}

func (this *builder) WithDescription(description string) *builder {
	this.priority.description = description
	return this
}

func (this *builder) WithCreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		this.fields = append(this.fields, messages.PriorityCreatedAt)
		this.errorMessages = append(this.errorMessages, messages.PriorityCreatedAtCannotBeEmpty)
		return this
	}
	this.priority.createdAt = createdAt
	return this
}

func (this *builder) WithUpdatedAt(updatedAt time.Time) *builder {
	this.priority.updatedAt = updatedAt
	return this
}

func (this *builder) Build() (*priority, errors.Error) {
	if len(this.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(this.errorMessages, map[string]interface{}{
			messages.Fields: this.fields})
	}
	return this.priority, nil
}
