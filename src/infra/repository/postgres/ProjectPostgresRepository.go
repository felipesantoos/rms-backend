package postgres

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/project"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
	"rms-backend/src/infra/repository/postgres/database"
)

type projectPostgresRepository struct{}

func NewProjectPostgresRepository() repository.IProjectRepository {
	return &projectPostgresRepository{}
}

func (*projectPostgresRepository) Create(projectObject project.Project) (*uuid.UUID, errors.Error) {
	transaction, err := database.BeginTransaction()
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	defer transaction.CloseConn()
	id, err := txQueryRowReturningID(
		transaction,
		database.Project().Command().Create(),
		projectObject.Name(),
		projectObject.Alias(),
		projectObject.Description(),
		projectObject.IsActive(),
	)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	projectID, conversionError := uuid.Parse(id)
	if conversionError != nil {
		logger.LogNativeError(conversionError)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	err = transaction.Commit()
	if err != nil {
		logger.LogCustomError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	return &projectID, nil
}
