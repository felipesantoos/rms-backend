package database

// Attributes

const (
	KnowledgeAreaID   = "knowledge_area_id"
	KnowledgeAreaName = "knowledge_area_name"
)

// Table

type KnowledgeAreaBuilder interface {
	Command() KnowledgeAreaCommandBuilder
	Query() KnowledgeAreaQueryBuilder
}

type knowledgeAreaBuilder struct{}

func KnowledgeArea() KnowledgeAreaBuilder {
	return knowledgeAreaBuilder{}
}

// Command

type KnowledgeAreaCommandBuilder interface {
	Create() string
	Update() string
	Delete() string
}

type knowledgeAreaCommandBuilder struct{}

func (knowledgeAreaBuilder) Command() KnowledgeAreaCommandBuilder {
	return knowledgeAreaCommandBuilder{}
}

func (knowledgeAreaCommandBuilder) Create() string {
	return empty
}

func (knowledgeAreaCommandBuilder) Update() string {
	return empty
}

func (knowledgeAreaCommandBuilder) Delete() string {
	return empty
}

// Query

type KnowledgeAreaQueryBuilder interface {
	All() string
	ByID() string
}

type knowledgeAreaQueryBuilder struct{}

func (knowledgeAreaBuilder) Query() KnowledgeAreaQueryBuilder {
	return knowledgeAreaQueryBuilder{}
}

func (knowledgeAreaQueryBuilder) All() string {
	return `
		SELECT 
		    knowledge_area.id AS knowledge_area_id, 
		    knowledge_area.name AS knowledge_area_name 
		FROM knowledge_area
	`
}

func (knowledgeAreaQueryBuilder) ByID() string {
	return empty
}
