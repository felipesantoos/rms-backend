package dicontainer

import (
	"rms-backend/src/core/interfaces/usecases"
	"rms-backend/src/core/services"
	"rms-backend/src/infra/repository/postgres"
)

func ResourcesUseCase() usecases.ResourcesUseCase {
	repo := postgres.NewResourcesPostgresAdapter()
	return services.NewResourcesService(repo)
}
