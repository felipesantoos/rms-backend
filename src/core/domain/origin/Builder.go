package origin

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
	origin        *origin
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		origin:        &origin{},
	}
}

func (this *builder) WithID(id uuid.UUID) *builder {
	if id.ID() == 0 {
		this.fields = append(this.fields, messages.OriginID)
		this.errorMessages = append(this.errorMessages, messages.OriginIDCannotBeEmpty)
		return this
	}
	this.origin.id = id
	return this
}

func (this *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		this.fields = append(this.fields, messages.OriginName)
		this.errorMessages = append(this.errorMessages, messages.OriginNameCannotBeEmpty)
		return this
	}
	this.origin.name = name
	return this
}

func (this *builder) WithDescription(description string) *builder {
	this.origin.description = description
	return this
}

func (this *builder) WithCreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		this.fields = append(this.fields, messages.OriginCreatedAt)
		this.errorMessages = append(this.errorMessages, messages.OriginCreatedAtCannotBeEmpty)
		return this
	}
	this.origin.createdAt = createdAt
	return this
}

func (this *builder) WithUpdatedAt(updatedAt time.Time) *builder {
	this.origin.updatedAt = updatedAt
	return this
}

func (this *builder) Build() (*origin, errors.Error) {
	if len(this.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(this.errorMessages, map[string]interface{}{
			messages.Fields: this.fields})
	}
	return this.origin, nil
}
