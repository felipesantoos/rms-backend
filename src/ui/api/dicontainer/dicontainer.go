package dicontainer

import (
	"rms-backend/src/core/interfaces/primary"
	"rms-backend/src/core/services"
	"rms-backend/src/infra/repository/postgres"
)

func ProjectServices() primary.IProjectServices {
	projectRepository := postgres.NewProjectPostgresRepository()
	return services.NewProjectServices(projectRepository)
}
