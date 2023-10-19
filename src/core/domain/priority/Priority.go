package priority

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"time"
)

var _ Priority = &priority{}

type Priority interface {
	ID() uuid.UUID
	Level() string
	Description() string
	CreatedAt() time.Time
	UpdatedAt() time.Time

	SetID(uuid.UUID) errors.Error
	SetLevel(string)
	SetDescription(string)
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
}

type priority struct {
	id          uuid.UUID
	level       string
	description string
	createdAt   time.Time
	updatedAt   time.Time
}

func (this *priority) ID() uuid.UUID {
	return this.id
}

func (this *priority) Level() string {
	return this.level
}

func (this *priority) Description() string {
	return this.description
}

func (this *priority) CreatedAt() time.Time {
	return this.createdAt
}

func (this *priority) UpdatedAt() time.Time {
	return this.updatedAt
}

func (this *priority) SetID(id uuid.UUID) errors.Error {
	this.id = id
	return nil
}

func (this *priority) SetLevel(level string) {
	this.level = level
}

func (this *priority) SetDescription(description string) {
	this.description = description
}

func (this *priority) SetCreatedAt(createdAt time.Time) {
	this.createdAt = createdAt
}

func (this *priority) SetUpdatedAt(updatedAt time.Time) {
	this.updatedAt = updatedAt
}
