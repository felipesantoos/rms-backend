package postgres

const (
	TruncateDemandStatus   = "TRUNCATE demand_status RESTART IDENTITY CASCADE"
	InsertIntoDemandStatus = "INSERT INTO demand_status (id, name, description, code) VALUES ('%v', '%v', '%v', '%v')"
)
