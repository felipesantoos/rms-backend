package project

import (
	"github.com/google/uuid"
	"time"
)

var _ Project = &project{}

type Project interface {
	ID() uuid.UUID
	Name() string
	Alias() string
	Description() string
	IsActive() bool
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() time.Time

	SetID(uuid.UUID)
	SetName(string)
	SetAlias(string)
	SetDescription(string)
	SetIsActive(bool)
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
	SetDeletedAt(time.Time)
}

type project struct {
	id          uuid.UUID
	name        string
	alias       string
	description string
	isActive    bool
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   time.Time
}

func (this *project) ID() uuid.UUID {
	return this.id
}

func (this *project) Name() string {
	return this.name
}

func (this *project) Alias() string {
	return this.alias
}

func (this *project) Description() string {
	return this.description
}

func (this *project) IsActive() bool {
	return this.isActive
}

func (this *project) CreatedAt() time.Time {
	return this.createdAt
}

func (this *project) UpdatedAt() time.Time {
	return this.updatedAt
}

func (this *project) DeletedAt() time.Time {
	return this.deletedAt
}

func (this *project) SetID(id uuid.UUID) {
	this.id = id
}

func (this *project) SetName(name string) {
	this.name = name
}

func (this *project) SetAlias(alias string) {
	this.alias = alias
}

func (this *project) SetDescription(description string) {
	this.description = description
}

func (this *project) SetIsActive(isActive bool) {
	this.isActive = isActive
}

func (this *project) SetCreatedAt(createdAt time.Time) {
	this.createdAt = createdAt
}

func (this *project) SetUpdatedAt(updatedAt time.Time) {
	this.updatedAt = updatedAt
}

func (this *project) SetDeletedAt(deletedAt time.Time) {
	this.deletedAt = deletedAt
}