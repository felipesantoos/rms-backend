package postgres

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/projectContainsUser"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
	"rms-backend/src/infra/repository/postgres/database"
	"rms-backend/src/infra/repository/postgres/queryObject"
)

type projectContainsUserPostgresRepository struct{}

func NewProjectContainsUserPostgresRepository() repository.IProjectContainsUserRepository {
	return &projectContainsUserPostgresRepository{}
}

func (*projectContainsUserPostgresRepository) Create(projectID uuid.UUID, projectContainsUserObject projectContainsUser.ProjectContainsUser) errors.Error {
	transaction, err := database.BeginTransaction()
	if err != nil {
		return logger.LogCustomError(err)
	}
	defer transaction.CloseConn()
	err = defaultTxExecQuery(
		transaction,
		database.ProjectContainsUser().Command().Create(),
		projectID,
		projectContainsUserObject.UserID(),
	)
	if err != nil {
		return logger.LogCustomError(err)
	}

	err = transaction.Commit()
	if err != nil {
		logger.LogCustomError(err)
		return logger.LogCustomError(errors.NewUnexpected())
	}
	return nil
}

func (this *projectContainsUserPostgresRepository) List(projectID uuid.UUID) ([]projectContainsUser.ProjectContainsUser, errors.Error) {
	rows, err := database.Queryx(database.ProjectContainsUser().Query().All(), projectID)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	projectContainsUsers := make([]projectContainsUser.ProjectContainsUser, 0)
	for rows.Next() {
		serializedProjectContainsUsers := map[string]interface{}{}
		rows.MapScan(serializedProjectContainsUsers)
		projectContainsUserObject, err := queryObject.NewProjectContainsUserFromMapRows(serializedProjectContainsUsers)
		if err != nil {
			return nil, logger.LogCustomError(err)
		}
		projectContainsUsers = append(projectContainsUsers, projectContainsUserObject)
	}
	return projectContainsUsers, nil
}

func (this *projectContainsUserPostgresRepository) Delete(projectID, userID uuid.UUID) errors.Error {
	err := defaultExecQuery(database.ProjectContainsUser().Command().Delete(), projectID, userID)
	if err != nil {
		return logger.LogCustomError(err)
	}
	return nil
}

func (this *projectContainsUserPostgresRepository) ReactivateCollaborator(projectID, userID uuid.UUID) errors.Error {
	err := defaultExecQuery(database.ProjectContainsUser().Command().ReactivateCollaborator(), projectID, userID)
	if err != nil {
		return logger.LogCustomError(err)
	}
	return nil
}

func (this *projectContainsUserPostgresRepository) IsCollaborator(projectID, userID uuid.UUID) (bool, errors.Error) {
	row, err := database.Queryx(database.ProjectContainsUser().Query().IsCollaborator(), projectID, userID)
	if err != nil {
		return false, logger.LogCustomError(err)
	}

	for !row.Next() {
		return false, nil
	}

	return true, nil
}
