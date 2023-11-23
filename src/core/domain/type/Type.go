package _type

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"time"
)

var _ Type = &_type{}

type Type interface {
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

type _type struct {
	id          uuid.UUID
	name        string
	description string
	createdAt   time.Time
	updatedAt   time.Time
}

func (this *_type) ID() uuid.UUID {
	return this.id
}

func (this *_type) Name() string {
	return this.name
}

func (this *_type) Description() string {
	return this.description
}

func (this *_type) CreatedAt() time.Time {
	return this.createdAt
}

func (this *_type) UpdatedAt() time.Time {
	return this.updatedAt
}

func (this *_type) SetID(id uuid.UUID) errors.Error {
	this.id = id
	return nil
}

func (this *_type) SetName(name string) {
	this.name = name
}

func (this *_type) SetDescription(description string) {
	this.description = description
}

func (this *_type) SetCreatedAt(createdAt time.Time) {
	this.createdAt = createdAt
}

func (this *_type) SetUpdatedAt(updatedAt time.Time) {
	this.updatedAt = updatedAt
}
