package projectContainsUser

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/messages"
	"time"
)

var _ ProjectContainsUser = &projectContainsUser{}

type ProjectContainsUser interface {
	UserID() uuid.UUID
	UserEmail() string
	UserFirstName() string
	UserLastName() string
	ProjectID() uuid.UUID
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() *time.Time

	SetUserID(uuid.UUID) errors.Error
	SetUserEmail(string)
	SetUserFirstName(string)
	SetUserLastName(string)
	SetProjectID(uuid.UUID) errors.Error
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
	SetDeletedAt(*time.Time)
}

type projectContainsUser struct {
	userId        uuid.UUID
	userEmail     string
	userFirstName string
	userLastName  string
	projectId     uuid.UUID
	createdAt     time.Time
	updatedAt     time.Time
	deletedAt     *time.Time
}

func (this *projectContainsUser) UserID() uuid.UUID {
	return this.userId
}

func (this *projectContainsUser) UserEmail() string {
	return this.userEmail
}

func (this *projectContainsUser) UserFirstName() string {
	return this.userFirstName
}

func (this *projectContainsUser) UserLastName() string {
	return this.userLastName
}

func (this *projectContainsUser) ProjectID() uuid.UUID {
	return this.projectId
}

func (this *projectContainsUser) CreatedAt() time.Time {
	return this.createdAt
}

func (this *projectContainsUser) UpdatedAt() time.Time {
	return this.updatedAt
}

func (this *projectContainsUser) DeletedAt() *time.Time {
	return this.deletedAt
}

func (this *projectContainsUser) SetUserID(userId uuid.UUID) errors.Error {
	if userId.ID() == 0 {
		return errors.NewValidationFromString(messages.UserIDCannotBeEmpty)
	}
	this.userId = userId
	return nil
}

func (this *projectContainsUser) SetUserEmail(email string) {
	this.userEmail = email
}

func (this *projectContainsUser) SetUserFirstName(firstName string) {
	this.userFirstName = firstName
}

func (this *projectContainsUser) SetUserLastName(lastName string) {
	this.userLastName = lastName
}

func (this *projectContainsUser) SetProjectID(projectId uuid.UUID) errors.Error {
	if projectId.ID() == 0 {
		return errors.NewValidationFromString(messages.ProjectIDCannotBeEmpty)
	}
	this.projectId = projectId
	return nil
}

func (this *projectContainsUser) SetCreatedAt(createdAt time.Time) {
	this.createdAt = createdAt
}

func (this *projectContainsUser) SetUpdatedAt(updatedAt time.Time) {
	this.createdAt = updatedAt
}

func (this *projectContainsUser) SetDeletedAt(deletedAt *time.Time) {
	this.deletedAt = deletedAt
}
