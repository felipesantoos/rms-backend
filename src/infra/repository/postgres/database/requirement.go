package database

// Attributes

const (
	RequirementID          = "requirement_id"
	RequirementCode        = "requirement_code"
	RequirementTitle       = "requirement_title"
	RequirementDescription = "requirement_description"
	RequirementUserStory   = "requirement_user_story"
	RequirementTypeID      = "requirement_type_id"
	RequirementOriginID    = "requirement_origin_id"
	RequirementPriorityID  = "requirement_priority_id"
	RequirementProjectID   = "requirement_project_id"
	RequirementCreatedAt   = "requirement_created_at"
	RequirementUpdatedAt   = "requirement_updated_at"
	RequirementDeletedAt   = "requirement_deleted_at"
)

// Table

type IRequirement interface {
	Command() IRequirementCommand
	Query() IRequirementQuery
}

type requirement struct{}

func Requirement() IRequirement {
	return requirement{}
}

// Command

type IRequirementCommand interface {
	Create() string
	Update() string
	Delete() string
}

type requirementCommand struct{}

func (requirement) Command() IRequirementCommand {
	return requirementCommand{}
}

func (requirementCommand) Create() string {
	return `
		INSERT INTO requirement (code, title, description, user_story, type_id, origin_id, priority_id, project_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`
}

func (requirementCommand) Update() string {
	return `
		UPDATE requirement SET
		   title = $1, 
		   description = $2, 
		   user_story = $3, 
		   type_id = $4, 
		   origin_id = $5, 
		   priority_id = $6, 
		   project_id = $7,
		   updated_at = NOW()
		WHERE id = $8
	`
}

func (requirementCommand) Delete() string {
	return `
		UPDATE requirement SET
		   deleted_at = NOW(),
		   updated_at = NOW()
		WHERE id = $1
	`
}

// Query

type IRequirementQuery interface {
	All() string
	ByID() string
	LastByProject() string
}

type requirementQuery struct{}

func (requirement) Query() IRequirementQuery {
	return requirementQuery{}
}

func (requirementQuery) All() string {
	return `
		SELECT
		    id AS requirement_id,
			code AS requirement_code,
			title AS requirement_title,
			description AS requirement_description,
			user_story AS requirement_user_story,
			type_id AS requirement_type_id,
			origin_id AS requirement_origin_id,
			priority_id AS requirement_priority_id,
			project_id AS requirement_project_id,
			created_at AS requirement_created_at,
			updated_at AS requirement_updated_at
		FROM requirement
		WHERE deleted_at IS NULL
			AND project_id = COALESCE($1, project_id)
		ORDER BY created_at DESC
	`
}

func (requirementQuery) ByID() string {
	return `
		SELECT
		    id AS requirement_id,
			code AS requirement_code,
			title AS requirement_title,
			description AS requirement_description,
			user_story AS requirement_user_story,
			type_id AS requirement_type_id,
			origin_id AS requirement_origin_id,
			priority_id AS requirement_priority_id,
			project_id AS requirement_project_id,
			created_at AS requirement_created_at,
			updated_at AS requirement_updated_at
		FROM requirement
		WHERE id = $1 AND deleted_at IS NULL
	`
}

func (requirementQuery) LastByProject() string {
	return `
		SELECT
		    id AS requirement_id,
			code AS requirement_code,
			title AS requirement_title,
			description AS requirement_description,
			user_story AS requirement_user_story,
			type_id AS requirement_type_id,
			origin_id AS requirement_origin_id,
			priority_id AS requirement_priority_id,
			project_id AS requirement_project_id,
			created_at AS requirement_created_at,
			updated_at AS requirement_updated_at
		FROM requirement
		WHERE project_id = $1 AND deleted_at IS NULL
		ORDER BY code DESC
		LIMIT 1
	`
}
