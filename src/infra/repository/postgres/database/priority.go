package database

// Attributes

const (
	PriorityID          = "priority_id"
	PriorityLevel       = "priority_level"
	PriorityDescription = "priority_description"
	PriorityCreatedAt   = "priority_created_at"
	PriorityUpdatedAt   = "priority_updated_at"
)

// Table

type IPriority interface {
	Command() IPriorityCommand
	Query() IPriorityQuery
}

type priority struct{}

func Priority() IPriority {
	return priority{}
}

// Command

type IPriorityCommand interface {
	Create() string
	Update() string
	Delete() string
}

type priorityCommand struct{}

func (priority) Command() IPriorityCommand {
	return priorityCommand{}
}

func (priorityCommand) Create() string {
	return empty
}

func (priorityCommand) Update() string {
	return empty
}

func (priorityCommand) Delete() string {
	return empty
}

// Query

type IPriorityQuery interface {
	All() string
	ByID() string
}

type priorityQuery struct{}

func (priority) Query() IPriorityQuery {
	return priorityQuery{}
}

func (priorityQuery) All() string {
	return `
		SELECT 
		    priority.id AS priority_id,
			priority.level AS priority_level,
			priority.description AS priority_description,
			priority.created_at AS priority_created_at,
			priority.updated_at AS priority_updated_at
		FROM priority
	`
}

func (priorityQuery) ByID() string {
	return empty
}
