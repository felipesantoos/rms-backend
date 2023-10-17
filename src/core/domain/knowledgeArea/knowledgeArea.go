package knowledgeArea

import (
	"strings"

	"github.com/google/uuid"
)

type KnowledgeArea interface {
	ID() uuid.UUID
	Name() string

	SetID(uuid.UUID)
	SetStringID(string) error
	SetName(string)
}

var _ KnowledgeArea = &knowledgeArea{}

type knowledgeArea struct {
	id   uuid.UUID
	name string
}

func (this *knowledgeArea) ID() uuid.UUID {
	return this.id
}

func (this *knowledgeArea) Name() string {
	return strings.ToUpper(this.name)
}

func (this *knowledgeArea) SetID(id uuid.UUID) {
	this.id = id
}

func (this *knowledgeArea) SetStringID(id string) error {
	if convertedID, err := uuid.Parse(id); err != nil {
		return err
	} else {
		this.id = convertedID
	}
	return nil
}

func (this *knowledgeArea) SetName(name string) {
	this.name = name
}
