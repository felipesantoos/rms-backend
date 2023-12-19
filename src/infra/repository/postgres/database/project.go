package database

// Attributes

const (
	ProjectID                 = "project_id"
	ProjectName               = "project_name"
	ProjectAlias              = "project_alias"
	ProjectDescription        = "project_description"
	ProjectIsActive           = "project_is_active"
	ProjectCreatedAt          = "project_created_at"
	ProjectUpdatedAt          = "project_updated_at"
	ProjectDeletedAt          = "project_deleted_at"
	ProjectCreatedByUserEmail = "project_created_by_user_email"
)

// Table

type IProject interface {
	Command() IProjectCommand
	Query() IProjectQuery
}

type project struct{}

func Project() IProject {
	return project{}
}

// Command

type IProjectCommand interface {
	Create() string
	Update() string
	Delete() string
}

type projectCommand struct{}

func (project) Command() IProjectCommand {
	return projectCommand{}
}

func (projectCommand) Create() string {
	return `
		INSERT INTO project (name, alias, description, is_active, created_by) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
}

func (projectCommand) Update() string {
	return `
		UPDATE project SET
		   name = $1,
		   alias = $2,
		   description = $3,
		   is_active = $4,
		   updated_at = NOW()
		WHERE id = $5
	`
}

func (projectCommand) Delete() string {
	return `
		UPDATE project SET
		   is_active = false,
		   deleted_at = NOW(),
		   updated_at = NOW()
		WHERE id = $1
	`
}

// Query

type IProjectQuery interface {
	All() string
	ByID() string
}

type projectQuery struct{}

func (project) Query() IProjectQuery {
	return projectQuery{}
}

func (projectQuery) All() string {
	return `
		SELECT
    		project.id AS project_id,
    		project.name AS project_name,
    		project.alias AS project_alias,
    		project.description AS project_description,
    		project.is_active AS project_is_active,
    		project.created_at AS project_created_at,
    		project.updated_at AS project_updated_at,
    		project.deleted_at AS project_deleted_at,
    		(SELECT account.email FROM account WHERE account.id = project.created_by) AS project_created_by_user_email
		FROM project
         	INNER JOIN project_contains_user ON project_contains_user.project_id = project.id
         	INNER JOIN account ON account.id = project_contains_user.user_id
		WHERE project.deleted_at IS NULL AND account.id = $1;
	`
}

func (projectQuery) ByID() string {
	return `
		SELECT 
		    project.id AS project_id,
			project.name AS project_name,
			project.alias AS project_alias,
			project.description AS project_description,
			project.is_active AS project_is_active,
			project.created_at AS project_created_at,
			project.updated_at AS project_updated_at,
			project.deleted_at AS project_deleted_at,
			(SELECT account.email FROM account WHERE account.id = project.created_by) AS project_created_by_user_email
		FROM project
		WHERE project.id = $1
			AND deleted_at IS NULL
	`
}
