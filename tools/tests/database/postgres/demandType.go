package postgres

const (
	TruncateDemandType   = "TRUNCATE demand_type RESTART IDENTITY CASCADE"
	InsertIntoDemandType = "INSERT INTO demand_type (id, name, description, code) VALUES ('%v', '%v', '%v', '%v')"
)
