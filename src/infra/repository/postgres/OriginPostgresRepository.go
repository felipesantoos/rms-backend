package postgres

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/origin"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
	"rms-backend/src/infra/repository/postgres/database"
	"rms-backend/src/infra/repository/postgres/queryObject"
)

type originPostgresRepository struct{}

func NewOriginPostgresRepository() repository.IOriginRepository {
	return &originPostgresRepository{}
}

func (this *originPostgresRepository) List() ([]origin.Origin, errors.Error) {
	rows, err := database.Queryx(database.Origin().Query().All())
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	origins := make([]origin.Origin, 0)
	for rows.Next() {
		serializedOrigin := map[string]interface{}{}
		rows.MapScan(serializedOrigin)
		originObject, err := queryObject.NewOriginFromMapRows(serializedOrigin)
		if err != nil {
			return nil, logger.LogCustomError(err)
		}
		origins = append(origins, originObject)
	}
	return origins, nil
}
