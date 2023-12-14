package database

// Attributes

const (
	ProjectContainsUserCreatedAt = "project_contains_user_created_at"
	ProjectContainsUserUpdatedAt = "project_contains_user_updated_at"
	ProjectContainsUserDeletedAt = "project_contains_user_deleted_at"
)

// Table

type IProjectContainsUser interface {
	Command() IProjectContainsUserCommand
	Query() IProjectContainsUserQuery
}

type projectContainsUser struct{}

func ProjectContainsUser() IProjectContainsUser {
	return projectContainsUser{}
}

// Command

type IProjectContainsUserCommand interface {
	Create() string
	Delete() string
	ReactivateCollaborator() string
}

type projectContainsUserCommand struct{}

func (projectContainsUser) Command() IProjectContainsUserCommand {
	return projectContainsUserCommand{}
}

func (projectContainsUserCommand) Create() string {
	return `
		INSERT INTO project_contains_user (project_id, user_id) 
		VALUES ($1, $2)
	`
}

func (projectContainsUserCommand) Delete() string {
	return `
		UPDATE project_contains_user SET
		   updated_at = NOW(),
		   deleted_at = NOW()
		WHERE project_id = $1
		AND user_id = $2
	`
}

func (projectContainsUserCommand) ReactivateCollaborator() string {
	return `
		UPDATE project_contains_user SET
			deleted_at = NULL,
			updated_at = NOW()
		WHERE project_id = $1
		AND user_id = $2
	`
}

// Query

type IProjectContainsUserQuery interface {
	All() string
	InverseAll() string
	IsCollaborator() string
}

type projectContainsUserQuery struct{}

func (projectContainsUser) Query() IProjectContainsUserQuery {
	return projectContainsUserQuery{}
}

func (projectContainsUserQuery) All() string {
	return `
		SELECT 
		    pcu.user_id AS user_id,
		    u.email AS user_email,
			u.first_name AS user_first_name,
			u.last_name AS user_last_name,
			pcu.created_at AS project_contains_user_created_at,
			pcu.updated_at AS project_contains_user_updated_at,
			pcu.deleted_at AS project_contains_user_deleted_at
		FROM project_contains_user pcu
		INNER JOIN account u ON pcu.user_id = u.id
		WHERE pcu.deleted_at IS NULL
		AND u.deleted_at IS NULL
		AND pcu.project_id = $1
		ORDER BY u.first_name
	`
}

func (projectContainsUserQuery) InverseAll() string {
	return `
		SELECT
    		u.id AS user_id,
    		u.email AS user_email,
    		u.first_name AS user_first_name,
    		u.last_name AS user_last_name,
    		pcu.created_at AS project_contains_user_created_at,
    		pcu.updated_at AS project_contains_user_updated_at,
    		pcu.deleted_at AS project_contains_user_deleted_at
		FROM account u
 		LEFT JOIN project_contains_user pcu ON u.id = pcu.user_id
	    AND pcu.deleted_at IS NULL
 		AND pcu.project_id = $1
		WHERE u.deleted_at IS NULL
		AND pcu.user_id IS NULL
		ORDER BY u.first_name
	`
}

func (projectContainsUserQuery) IsCollaborator() string {
	return `
		SELECT 
		    pcu.user_id AS user_id,
		    u.email AS user_email,
			u.first_name AS user_first_name,
			u.last_name AS user_last_name,
			pcu.created_at AS project_contains_user_created_at,
			pcu.updated_at AS project_contains_user_updated_at,
			pcu.deleted_at AS project_contains_user_deleted_at
		FROM project_contains_user pcu
		INNER JOIN account u ON pcu.user_id = u.id
		WHERE u.deleted_at IS NULL
		AND pcu.project_id = $1
		AND pcu.user_id = $2
	`
}
