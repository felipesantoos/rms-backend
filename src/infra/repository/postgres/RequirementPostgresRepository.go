package postgres

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/requirement"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/messages"
	"rms-backend/src/core/services/filters"
	"rms-backend/src/infra/repository/postgres/database"
	"rms-backend/src/infra/repository/postgres/queryObject"
)

type requirementPostgresRepository struct{}

func NewRequirementPostgresRepository() repository.IRequirementRepository {
	return &requirementPostgresRepository{}
}

func (*requirementPostgresRepository) Create(requirementObject requirement.Requirement) (*uuid.UUID, errors.Error) {
	transaction, err := database.BeginTransaction()
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	defer transaction.CloseConn()
	id, err := txQueryRowReturningID(
		transaction,
		database.Requirement().Command().Create(),
		requirementObject.Code(),
		requirementObject.Title(),
		requirementObject.Description(),
		requirementObject.UserStory(),
		requirementObject.TypeID(),
		requirementObject.OriginID(),
		requirementObject.PriorityID(),
		requirementObject.ProjectID(),
	)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	requirementID, conversionError := uuid.Parse(id)
	if conversionError != nil {
		logger.LogNativeError(conversionError)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	err = transaction.Commit()
	if err != nil {
		logger.LogCustomError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	return &requirementID, nil
}

func (this *requirementPostgresRepository) List(requirementFilters filters.RequirementFilters) ([]requirement.Requirement, errors.Error) {
	rows, err := database.Queryx(database.Requirement().Query().All(), requirementFilters.ProjectID)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	requirements := make([]requirement.Requirement, 0)
	for rows.Next() {
		serializedRequirement := map[string]interface{}{}
		rows.MapScan(serializedRequirement)
		requirementObject, err := queryObject.NewRequirementFromMapRows(serializedRequirement)
		if err != nil {
			return nil, logger.LogCustomError(err)
		}
		requirements = append(requirements, requirementObject)
	}
	return requirements, nil
}

func (this *requirementPostgresRepository) Get(id uuid.UUID) (requirement.Requirement, errors.Error) {
	row, err := database.Queryx(database.Requirement().Query().ByID(), id)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	for !row.Next() {
		return nil, logger.LogCustomError(errors.NewFromString(messages.RequirementNotFoundErrorMessage))
	}
	serializedRequirement := map[string]interface{}{}
	row.MapScan(serializedRequirement)
	requirementObject, err := queryObject.NewRequirementFromMapRows(serializedRequirement)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return requirementObject, nil
}

func (this *requirementPostgresRepository) GetLastCodeByProject(projectID uuid.UUID) (*int, errors.Error) {
	row, err := database.Queryx(database.Requirement().Query().LastByProject(), projectID)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	for !row.Next() {
		return nil, logger.LogCustomError(errors.NewFromString(messages.RequirementNotFoundErrorMessage))
	}
	serializedRequirement := map[string]interface{}{}
	row.MapScan(serializedRequirement)
	requirementObject, err := queryObject.NewRequirementFromMapRows(serializedRequirement)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	code := requirementObject.Code()
	return &code, nil
}

func (this *requirementPostgresRepository) Update(requirementObject requirement.Requirement) errors.Error {
	err := defaultExecQuery(
		database.Requirement().Command().Update(),
		requirementObject.Title(),
		requirementObject.Description(),
		requirementObject.UserStory(),
		requirementObject.TypeID(),
		requirementObject.OriginID(),
		requirementObject.PriorityID(),
		requirementObject.ProjectID(),
		requirementObject.ID(),
	)
	if err != nil {
		return logger.LogCustomError(err)
	}
	return nil
}

func (this *requirementPostgresRepository) Delete(id uuid.UUID) errors.Error {
	err := defaultExecQuery(database.Requirement().Command().Delete(), id)
	if err != nil {
		return logger.LogCustomError(err)
	}
	return nil
	return nil
}
