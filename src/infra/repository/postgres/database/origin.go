package database

// Attributes

const (
	OriginID          = "origin_id"
	OriginName        = "origin_name"
	OriginDescription = "origin_description"
	OriginCreatedAt   = "origin_created_at"
	OriginUpdatedAt   = "priority_updated_at"
)

// Table

type IOrigin interface {
	Command() IOriginCommand
	Query() IOriginQuery
}

type origin struct{}

func Origin() IOrigin {
	return origin{}
}

// Command

type IOriginCommand interface {
	Create() string
	Update() string
	Delete() string
}

type originCommand struct{}

func (origin) Command() IOriginCommand {
	return originCommand{}
}

func (originCommand) Create() string {
	return empty
}

func (originCommand) Update() string {
	return empty
}

func (originCommand) Delete() string {
	return empty
}

// Query

type IOriginQuery interface {
	All() string
	ByID() string
}

type originQuery struct{}

func (origin) Query() IOriginQuery {
	return originQuery{}
}

func (originQuery) All() string {
	return `
		SELECT 
		    origin.id AS origin_id,
			origin.name AS origin_name,
			origin.description AS origin_description,
			origin.created_at AS origin_created_at,
			origin.updated_at AS origin_updated_at
		FROM requirement_origin origin
	`
}

func (originQuery) ByID() string {
	return empty
}
