package database

// Attributes

const (
	UserID        = "user_id"
	UserEmail     = "user_email"
	UserFirstName = "user_first_name"
	UserLastName  = "user_last_name"
	UserIsActive  = "user_is_active"
	UserCreatedAt = "user_created_at"
	UserUpdatedAt = "user_updated_at"
	UserDeletedAt = "user_deleted_at"
)

// Table

type IUser interface {
	Command() IUserCommand
	Query() IUserQuery
}

type user struct{}

func User() IUser {
	return user{}
}

// Command

type IUserCommand interface {
	Create() string
	Delete() string
}

type userCommand struct{}

func (user) Command() IUserCommand {
	return userCommand{}
}

func (userCommand) Create() string {
	return `
		
	`
}

func (userCommand) Delete() string {
	return `
		
	`
}

// Query

type IUserQuery interface {
	All() string
	ByID() string
	ByEmail() string
}

type userQuery struct{}

func (user) Query() IUserQuery {
	return userQuery{}
}

func (userQuery) All() string {
	return `
		
	`
}

func (userQuery) ByID() string {
	return `
		SELECT
			u.id AS user_id,
			u.email AS user_email,
			u.first_name AS user_first_name,
			u.last_name AS user_last_name,
			u.is_active AS user_is_active,
			u.created_at AS user_created_at,
			u.updated_at AS user_updated_at,
			u.deleted_at AS user_deleted_at
		FROM account u
		WHERE u.deleted_at IS NULL
		AND u.id = $1
    `
}

func (userQuery) ByEmail() string {
	return `
		SELECT
			u.id AS user_id,
			u.email AS user_email,
			u.first_name AS user_first_name,
			u.last_name AS user_last_name,
			u.is_active AS user_is_active,
			u.created_at AS user_created_at,
			u.updated_at AS user_updated_at,
			u.deleted_at AS user_deleted_at
		FROM account u
		WHERE u.deleted_at IS NULL
		AND u.email = $1
    `
}
