package postgres

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/project"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/messages"
	"rms-backend/src/core/services/filters"
	"rms-backend/src/infra/repository/postgres/database"
	"rms-backend/src/infra/repository/postgres/queryObject"
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

func (this *projectPostgresRepository) List(projectFilters filters.ProjectFilters) ([]project.Project, errors.Error) {
	rows, err := database.Queryx(database.Project().Query().All(), *projectFilters.UserID)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	projects := make([]project.Project, 0)
	for rows.Next() {
		serializedProject := map[string]interface{}{}
		rows.MapScan(serializedProject)
		projectObject, err := queryObject.NewProjectFromMapRows(serializedProject)
		if err != nil {
			return nil, logger.LogCustomError(err)
		}
		projects = append(projects, projectObject)
	}
	return projects, nil
}

func (this *projectPostgresRepository) Get(id uuid.UUID) (project.Project, errors.Error) {
	row, err := database.Queryx(database.Project().Query().ByID(), id)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	for !row.Next() {
		return nil, logger.LogCustomError(errors.NewFromString(messages.ProjectNotFoundErrorMessage))
	}
	serializedProject := map[string]interface{}{}
	row.MapScan(serializedProject)
	projectObject, err := queryObject.NewProjectFromMapRows(serializedProject)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return projectObject, nil
}

func (this *projectPostgresRepository) Update(projectObject project.Project) errors.Error {
	err := defaultExecQuery(
		database.Project().Command().Update(),
		projectObject.Name(),
		projectObject.Alias(),
		projectObject.Description(),
		projectObject.IsActive(),
		projectObject.ID(),
	)
	if err != nil {
		return logger.LogCustomError(err)
	}
	return nil
}

func (this *projectPostgresRepository) Delete(id uuid.UUID) errors.Error {
	err := defaultExecQuery(database.Project().Command().Delete(), id)
	if err != nil {
		return logger.LogCustomError(err)
	}
	return nil
}
