package postgres

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/priority"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
	"rms-backend/src/infra/repository/postgres/database"
	"rms-backend/src/infra/repository/postgres/queryObject"
)

type priorityPostgresRepository struct{}

func NewPriorityPostgresRepository() repository.IPriorityRepository {
	return &priorityPostgresRepository{}
}

func (this *priorityPostgresRepository) List() ([]priority.Priority, errors.Error) {
	rows, err := database.Queryx(database.Priority().Query().All())
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	prioritys := make([]priority.Priority, 0)
	for rows.Next() {
		serializedPriority := map[string]interface{}{}
		rows.MapScan(serializedPriority)
		priorityObject, err := queryObject.NewPriorityFromMapRows(serializedPriority)
		if err != nil {
			return nil, logger.LogCustomError(err)
		}
		prioritys = append(prioritys, priorityObject)
	}
	return prioritys, nil
}
