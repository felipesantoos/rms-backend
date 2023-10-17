package knowledgeArea

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/messages"
	"strings"
)

type builder struct {
	errorMessages []string
	fields        []string
	knowledgeArea *knowledgeArea
}

func NewBuilder() *builder {
	return &builder{errorMessages: []string{}, fields: []string{}, knowledgeArea: &knowledgeArea{}}
}

func (this *builder) WithID(id uuid.UUID) *builder {
	if id.ID() == 0 {
		this.fields = append(this.fields, messages.KnowledgeAreaID)
		this.errorMessages = append(this.fields, messages.KnowledgeAreaIDCannotBeEmpty)
		return this
	}
	this.knowledgeArea.id = id
	return this
}

func (this *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		this.fields = append(this.fields, messages.KnowledgeAreaName)
		this.errorMessages = append(this.fields, messages.KnowledgeAreaNameCannotBeEmpty)
		return this
	}
	this.knowledgeArea.name = name
	return this
}

func (this *builder) Build() (*knowledgeArea, errors.Error) {
	if len(this.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(this.errorMessages, map[string]interface{}{
			messages.Fields: this.fields})
	}
	return this.knowledgeArea, nil
}
