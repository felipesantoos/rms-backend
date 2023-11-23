package origin

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"time"
)

var _ Origin = &origin{}

type Origin interface {
	ID() uuid.UUID
	Name() string
	Description() string
	CreatedAt() time.Time
	UpdatedAt() time.Time

	SetID(uuid.UUID) errors.Error
	SetName(string)
	SetDescription(string)
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
}

type origin struct {
	id          uuid.UUID
	name        string
	description string
	createdAt   time.Time
	updatedAt   time.Time
}

func (this *origin) ID() uuid.UUID {
	return this.id
}

func (this *origin) Name() string {
	return this.name
}

func (this *origin) Description() string {
	return this.description
}

func (this *origin) CreatedAt() time.Time {
	return this.createdAt
}

func (this *origin) UpdatedAt() time.Time {
	return this.updatedAt
}

func (this *origin) SetID(id uuid.UUID) errors.Error {
	this.id = id
	return nil
}

func (this *origin) SetName(name string) {
	this.name = name
}

func (this *origin) SetDescription(description string) {
	this.description = description
}

func (this *origin) SetCreatedAt(createdAt time.Time) {
	this.createdAt = createdAt
}

func (this *origin) SetUpdatedAt(updatedAt time.Time) {
	this.updatedAt = updatedAt
}
