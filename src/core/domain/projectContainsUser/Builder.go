package projectContainsUser

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/messages"
	"strings"
	"time"
)

type builder struct {
	fields              []string
	errorMessages       []string
	projectContainsUser *projectContainsUser
}

func NewBuilder() *builder {
	return &builder{
		fields:              []string{},
		errorMessages:       []string{},
		projectContainsUser: &projectContainsUser{},
	}
}

func (this *builder) WithUserID(userId uuid.UUID) *builder {
	if userId.ID() == 0 {
		this.fields = append(this.fields, messages.UserID)
		this.errorMessages = append(this.errorMessages, messages.UserIDCannotBeEmpty)
		return this
	}
	this.projectContainsUser.userId = userId
	return this
}

func (this *builder) WithUserEmail(email string) *builder {
	email = strings.TrimSpace(email)
	if len(email) == 0 {
		this.fields = append(this.fields, messages.UserEmail)
		this.errorMessages = append(this.errorMessages, messages.UserEmailCannotBeEmpty)
		return this
	}
	this.projectContainsUser.userEmail = email
	return this
}

func (this *builder) WithFirstName(firstName string) *builder {
	firstName = strings.TrimSpace(firstName)
	if len(firstName) == 0 {
		this.fields = append(this.fields, messages.UserFirstName)
		this.errorMessages = append(this.errorMessages, messages.UserFirstNameCannotBeEmpty)
		return this
	}
	this.projectContainsUser.userFirstName = firstName
	return this
}

func (this *builder) WithLastName(lastName string) *builder {
	lastName = strings.TrimSpace(lastName)
	if len(lastName) == 0 {
		this.fields = append(this.fields, messages.UserLastName)
		this.errorMessages = append(this.errorMessages, messages.UserLastNameCannotBeEmpty)
		return this
	}
	this.projectContainsUser.userLastName = lastName
	return this
}

func (this *builder) WithCreatedAt(createdAt *time.Time) *builder {
	this.projectContainsUser.createdAt = createdAt
	return this
}

func (this *builder) WithUpdatedAt(updatedAt *time.Time) *builder {
	this.projectContainsUser.updatedAt = updatedAt
	return this
}

func (this *builder) WithDeletedAt(deletedAt *time.Time) *builder {
	this.projectContainsUser.deletedAt = deletedAt
	return this
}

func (this *builder) Build() (*projectContainsUser, errors.Error) {
	if len(this.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(this.errorMessages, map[string]interface{}{
			messages.Fields: this.fields})
	}
	return this.projectContainsUser, nil
}
