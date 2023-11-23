package requirement

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/messages"
	"time"
)

var _ Requirement = &requirement{}

type Requirement interface {
	ID() uuid.UUID
	Code() int
	Title() string
	Description() string
	UserStory() string
	TypeID() uuid.UUID
	OriginID() uuid.UUID
	PriorityID() uuid.UUID
	ProjectID() uuid.UUID
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() *time.Time

	SetID(uuid.UUID) errors.Error
	SetCode(int)
	SetTitle(string)
	SetDescription(string)
	SetUserStory(string)
	SetTypeID(uuid.UUID) errors.Error
	SetOriginID(uuid.UUID) errors.Error
	SetPriorityID(uuid.UUID) errors.Error
	SetProjectID(uuid.UUID) errors.Error
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
	SetDeletedAt(*time.Time)
}

type requirement struct {
	id          uuid.UUID
	code        int
	title       string
	description string
	userStory   string
	typeID      uuid.UUID
	originID    uuid.UUID
	priorityID  uuid.UUID
	projectID   uuid.UUID
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   *time.Time
}

func (this *requirement) ID() uuid.UUID {
	return this.id
}

func (this *requirement) Code() int {
	return this.code
}

func (this *requirement) Title() string {
	return this.title
}

func (this *requirement) Description() string {
	return this.description
}

func (this *requirement) UserStory() string {
	return this.userStory
}

func (this *requirement) TypeID() uuid.UUID {
	return this.typeID
}

func (this *requirement) OriginID() uuid.UUID {
	return this.originID
}

func (this *requirement) PriorityID() uuid.UUID {
	return this.priorityID
}

func (this *requirement) ProjectID() uuid.UUID {
	return this.projectID
}

func (this *requirement) CreatedAt() time.Time {
	return this.createdAt
}

func (this *requirement) UpdatedAt() time.Time {
	return this.updatedAt
}

func (this *requirement) DeletedAt() *time.Time {
	return this.deletedAt
}

func (this *requirement) SetID(id uuid.UUID) errors.Error {
	if id.ID() == 0 {
		return errors.NewValidationFromString(messages.RequirementIDCannotBeEmpty)
	}
	this.id = id
	return nil
}

func (this *requirement) SetCode(code int) {
	this.code = code
}

func (this *requirement) SetTitle(title string) {
	this.title = title
}

func (this *requirement) SetDescription(description string) {
	this.description = description
}

func (this *requirement) SetUserStory(userStory string) {
	this.userStory = userStory
}

func (this *requirement) SetTypeID(typeID uuid.UUID) errors.Error {
	if typeID.ID() == 0 {
		return errors.NewValidationFromString(messages.RequirementTypeIDCannotBeEmpty)
	}
	this.typeID = typeID
	return nil
}

func (this *requirement) SetOriginID(originID uuid.UUID) errors.Error {
	if originID.ID() == 0 {
		return errors.NewValidationFromString(messages.RequirementOriginIDCannotBeEmpty)
	}
	this.originID = originID
	return nil
}

func (this *requirement) SetPriorityID(priorityID uuid.UUID) errors.Error {
	if priorityID.ID() == 0 {
		return errors.NewValidationFromString(messages.RequirementPriorityIDCannotBeEmpty)
	}
	this.priorityID = priorityID
	return nil
}

func (this *requirement) SetProjectID(projectID uuid.UUID) errors.Error {
	if projectID.ID() == 0 {
		return errors.NewValidationFromString(messages.RequirementProjectIDCannotBeEmpty)
	}
	this.projectID = projectID
	return nil
}

func (this *requirement) SetCreatedAt(createdAt time.Time) {
	this.createdAt = createdAt
}

func (this *requirement) SetUpdatedAt(updatedAt time.Time) {
	this.updatedAt = updatedAt
}

func (this *requirement) SetDeletedAt(deletedAt *time.Time) {
	this.deletedAt = deletedAt
}
