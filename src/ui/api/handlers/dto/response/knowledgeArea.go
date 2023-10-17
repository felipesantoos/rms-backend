package response

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/knowledgeArea"
)

type KnowledgeArea struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type knowledgeAreaBuilder struct{}

func KnowledgeAreaBuilder() *knowledgeAreaBuilder {
	return &knowledgeAreaBuilder{}
}

func (*knowledgeAreaBuilder) BuildFromDomain(data knowledgeArea.KnowledgeArea) KnowledgeArea {
	return KnowledgeArea{
		ID:   data.ID(),
		Name: data.Name(),
	}
}

func (builder *knowledgeAreaBuilder) BuildFromDomainList(list []knowledgeArea.KnowledgeArea) []KnowledgeArea {
	result := []KnowledgeArea{}
	for _, item := range list {
		result = append(result, builder.BuildFromDomain(item))
	}
	return result
}
