package queryObject

import (
	"fmt"
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/project"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/utils"
	"rms-backend/src/infra/repository/postgres/database"
	"strconv"
	"time"
)

func NewProjectFromMapRows(data map[string]interface{}) (project.Project, errors.Error) {
	var err error
	id, err := uuid.Parse(string(data[database.ProjectID].([]uint8)))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	var name = fmt.Sprint(data[database.ProjectName])
	var alias = fmt.Sprint(data[database.ProjectAlias])
	var description = fmt.Sprint(data[database.ProjectDescription])
	isActive, err := strconv.ParseBool(fmt.Sprint(data[database.ProjectIsActive]))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	nullableCreatedAt := utils.GetNullableValue[time.Time](data[database.ProjectCreatedAt])
	nullableUpdatedAt := utils.GetNullableValue[time.Time](data[database.ProjectUpdatedAt])
	var deletedAt = utils.GetNullableValue[time.Time](data[database.ProjectDeletedAt])
	var createdAt time.Time
	if nullableCreatedAt != nil {
		createdAt = *nullableCreatedAt
	}
	var updatedAt time.Time
	if nullableUpdatedAt != nil {
		updatedAt = *nullableUpdatedAt
	}
	projectObject, validationError := project.NewBuilder().WithID(id).WithName(name).WithAlias(alias).
		WithDescription(description).WithIsActive(isActive).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).
		WithDeletedAt(deletedAt).Build()
	if validationError != nil {
		return nil, logger.LogCustomError(validationError)
	}
	return projectObject, nil
}
