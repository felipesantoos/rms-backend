package _type

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
	_type         *_type
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		_type:         &_type{},
	}
}

func (this *builder) WithID(id uuid.UUID) *builder {
	if id.ID() == 0 {
		this.fields = append(this.fields, messages.TypeID)
		this.errorMessages = append(this.errorMessages, messages.TypeIDCannotBeEmpty)
		return this
	}
	this._type.id = id
	return this
}

func (this *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		this.fields = append(this.fields, messages.TypeName)
		this.errorMessages = append(this.errorMessages, messages.TypeNameCannotBeEmpty)
		return this
	}
	this._type.name = name
	return this
}

func (this *builder) WithDescription(description string) *builder {
	this._type.description = description
	return this
}

func (this *builder) WithCreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		this.fields = append(this.fields, messages.TypeCreatedAt)
		this.errorMessages = append(this.errorMessages, messages.TypeCreatedAtCannotBeEmpty)
		return this
	}
	this._type.createdAt = createdAt
	return this
}

func (this *builder) WithUpdatedAt(updatedAt time.Time) *builder {
	this._type.updatedAt = updatedAt
	return this
}

func (this *builder) Build() (*_type, errors.Error) {
	if len(this.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(this.errorMessages, map[string]interface{}{
			messages.Fields: this.fields})
	}
	return this._type, nil
}
