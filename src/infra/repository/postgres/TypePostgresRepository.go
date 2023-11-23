package postgres

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/type"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
	"rms-backend/src/infra/repository/postgres/database"
	"rms-backend/src/infra/repository/postgres/queryObject"
)

type typePostgresRepository struct{}

func NewTypePostgresRepository() repository.ITypeRepository {
	return &typePostgresRepository{}
}

func (this *typePostgresRepository) List() ([]_type.Type, errors.Error) {
	rows, err := database.Queryx(database.Type().Query().All())
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	types := make([]_type.Type, 0)
	for rows.Next() {
		serializedType := map[string]interface{}{}
		rows.MapScan(serializedType)
		typeObject, err := queryObject.NewTypeFromMapRows(serializedType)
		if err != nil {
			return nil, logger.LogCustomError(err)
		}
		types = append(types, typeObject)
	}
	return types, nil
}
