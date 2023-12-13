package user

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/messages"
	"time"
)

var _ User = &user{}

type User interface {
	ID() uuid.UUID
	Email() string
	FirstName() string
	LastName() string
	Password() string
	Salt() string
	IsActive() bool
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() *time.Time

	SetID(uuid.UUID) errors.Error
	SetEmail(string)
	SetFirstName(string)
	SetLastName(string)
	SetPassword(string)
	SetSalt(string)
	SetIsActive(bool)
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
	SetDeletedAt(*time.Time)
}

type user struct {
	id        uuid.UUID
	email     string
	firstName string
	lastName  string
	password  string
	salt      string
	isActive  bool
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func (this *user) ID() uuid.UUID {
	return this.id
}

func (this *user) Email() string {
	return this.email
}

func (this *user) FirstName() string {
	return this.firstName
}

func (this *user) LastName() string {
	return this.firstName
}

func (this *user) Password() string {
	return this.password
}

func (this *user) Salt() string {
	return this.salt
}

func (this *user) IsActive() bool {
	return this.isActive
}

func (this *user) CreatedAt() time.Time {
	return this.createdAt
}

func (this *user) UpdatedAt() time.Time {
	return this.updatedAt
}

func (this *user) DeletedAt() *time.Time {
	return this.deletedAt
}

func (this *user) SetID(id uuid.UUID) errors.Error {
	if id.ID() == 0 {
		return errors.NewValidationFromString(messages.UserIDCannotBeEmpty)
	}
	this.id = id
	return nil
}

func (this *user) SetEmail(email string) {
	this.email = email
}

func (this *user) SetFirstName(firstName string) {
	this.firstName = firstName
}

func (this *user) SetLastName(lastName string) {
	this.lastName = lastName
}

func (this *user) SetPassword(password string) {
	this.password = password
}

func (this *user) SetSalt(salt string) {
	this.salt = salt
}

func (this *user) SetIsActive(isActive bool) {
	this.isActive = isActive
}

func (this *user) SetCreatedAt(createdAt time.Time) {
	this.createdAt = createdAt
}

func (this *user) SetUpdatedAt(updatedAt time.Time) {
	this.updatedAt = updatedAt
}

func (this *user) SetDeletedAt(deletedAt *time.Time) {
	this.deletedAt = deletedAt
}
