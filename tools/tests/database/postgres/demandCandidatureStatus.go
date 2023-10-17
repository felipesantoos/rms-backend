package postgres

const (
	TruncateDemandCandidatureStatus   = "TRUNCATE demand_candidature_status RESTART IDENTITY CASCADE"
	InsertIntoDemandCandidatureStatus = "INSERT INTO demand_candidature_status (id, name, description, code) VALUES ('%v', '%v', '%v', '%v')"
)
