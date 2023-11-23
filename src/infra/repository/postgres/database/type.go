package database

// Attributes

const (
	TypeID          = "type_id"
	TypeName        = "type_name"
	TypeDescription = "type_description"
	TypeCreatedAt   = "type_created_at"
	TypeUpdatedAt   = "priority_updated_at"
)

// Table

type IType interface {
	Command() ITypeCommand
	Query() ITypeQuery
}

type _type struct{}

func Type() IType {
	return _type{}
}

// Command

type ITypeCommand interface {
	Create() string
	Update() string
	Delete() string
}

type typeCommand struct{}

func (_type) Command() ITypeCommand {
	return typeCommand{}
}

func (typeCommand) Create() string {
	return empty
}

func (typeCommand) Update() string {
	return empty
}

func (typeCommand) Delete() string {
	return empty
}

// Query

type ITypeQuery interface {
	All() string
	ByID() string
}

type typeQuery struct{}

func (_type) Query() ITypeQuery {
	return typeQuery{}
}

func (typeQuery) All() string {
	return `
		SELECT 
		    type.id AS type_id,
			type.name AS type_name,
			type.description AS type_description,
			type.created_at AS type_created_at,
			type.updated_at AS type_updated_at
		FROM requirement_type type
	`
}

func (typeQuery) ByID() string {
	return empty
}
