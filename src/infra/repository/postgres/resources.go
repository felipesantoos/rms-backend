package postgres

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/knowledgeArea"
	"rms-backend/src/core/interfaces/adapters"
	"rms-backend/src/core/logger"
	"rms-backend/src/infra/repository"
	"rms-backend/src/infra/repository/postgres/database"
	"rms-backend/src/infra/repository/postgres/queryObject"
)

type resourcesPostgresAdapter struct{}

func NewResourcesPostgresAdapter() adapters.ResourcesAdapter {
	return &resourcesPostgresAdapter{}
}

func (*resourcesPostgresAdapter) ListKnowledgeAreas() ([]knowledgeArea.KnowledgeArea, errors.Error) {
	rows, err := repository.Queryx(database.KnowledgeArea().Query().All())
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	var knowledgeAreas = make([]knowledgeArea.KnowledgeArea, 0)
	for rows.Next() {
		var serializedKnowledgeArea = map[string]interface{}{}
		rows.MapScan(serializedKnowledgeArea)
		_knowledgeArea, err := queryObject.NewKnowledgeAreaFromMapRows(serializedKnowledgeArea)
		if err != nil {
			logger.LogCustomError(err)
			return nil, logger.LogCustomError(errors.NewUnexpected())
		}
		knowledgeAreas = append(knowledgeAreas, _knowledgeArea)
	}
	return knowledgeAreas, nil
}
