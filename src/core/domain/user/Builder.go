package user

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
	user          *user
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		user:          &user{},
	}
}

func (this *builder) WithID(id uuid.UUID) *builder {
	if id.ID() == 0 {
		this.fields = append(this.fields, messages.UserID)
		this.errorMessages = append(this.errorMessages, messages.UserIDCannotBeEmpty)
		return this
	}
	this.user.id = id
	return this
}

func (this *builder) WithEmail(email string) *builder {
	email = strings.TrimSpace(email)
	if len(email) == 0 {
		this.fields = append(this.fields, messages.UserEmail)
		this.errorMessages = append(this.errorMessages, messages.UserEmailCannotBeEmpty)
		return this
	}
	this.user.email = email
	return this
}

func (this *builder) WithFirstName(firstName string) *builder {
	firstName = strings.TrimSpace(firstName)
	if len(firstName) == 0 {
		this.fields = append(this.fields, messages.UserFirstName)
		this.errorMessages = append(this.errorMessages, messages.UserFirstNameCannotBeEmpty)
		return this
	}
	this.user.firstName = firstName
	return this
}

func (this *builder) WithLastName(lastName string) *builder {
	lastName = strings.TrimSpace(lastName)
	if len(lastName) == 0 {
		this.fields = append(this.fields, messages.UserLastName)
		this.errorMessages = append(this.errorMessages, messages.UserLastNameCannotBeEmpty)
		return this
	}
	this.user.lastName = lastName
	return this
}

func (this *builder) WithPassword(password string) *builder {
	password = strings.TrimSpace(password)
	if len(password) == 0 {
		this.fields = append(this.fields, messages.UserSalt)
		this.errorMessages = append(this.errorMessages, messages.UserSaltCannotBeEmpty)
		return this
	}
	this.user.password = password
	return this
}

func (this *builder) WithSalt(salt string) *builder {
	salt = strings.TrimSpace(salt)
	if len(salt) == 0 {
		this.fields = append(this.fields, messages.UserSalt)
		this.errorMessages = append(this.errorMessages, messages.UserSaltCannotBeEmpty)
		return this
	}
	this.user.salt = salt
	return this
}

func (this *builder) WithIsActive(isActive bool) *builder {
	this.user.isActive = isActive
	return this
}

func (this *builder) WithCreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		this.fields = append(this.fields, messages.UserCreatedAt)
		this.errorMessages = append(this.errorMessages, messages.UserCreatedAtCannotBeEmpty)
		return this
	}
	this.user.createdAt = createdAt
	return this
}

func (this *builder) WithUpdatedAt(updatedAt time.Time) *builder {
	this.user.updatedAt = updatedAt
	return this
}

func (this *builder) WithDeletedAt(deletedAt *time.Time) *builder {
	this.user.deletedAt = deletedAt
	return this
}

func (this *builder) Build() (*user, errors.Error) {
	if len(this.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(this.errorMessages, map[string]interface{}{
			messages.Fields: this.fields})
	}
	return this.user, nil
}
