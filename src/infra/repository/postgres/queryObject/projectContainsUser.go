package queryObject

import (
	"fmt"
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/projectContainsUser"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/utils"
	"rms-backend/src/infra/repository/postgres/database"
	"time"
)

func NewProjectContainsUserFromMapRows(data map[string]interface{}) (projectContainsUser.ProjectContainsUser, errors.Error) {
	var err error
	userId, err := uuid.Parse(string(data[database.UserID].([]uint8)))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}

	var userEmail = fmt.Sprint(data[database.UserEmail])
	var userFirstName = fmt.Sprint(data[database.UserFirstName])
	var userLastName = fmt.Sprint(data[database.UserLastName])

	projectId, err := uuid.Parse(string(data[database.ProjectID].([]uint8)))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	nullableCreatedAt := utils.GetNullableValue[time.Time](data[database.ProjectContainsUserCreatedAt])
	nullableUpdatedAt := utils.GetNullableValue[time.Time](data[database.ProjectContainsUserUpdatedAt])
	var deletedAt = utils.GetNullableValue[time.Time](data[database.ProjectContainsUserDeletedAt])
	var createdAt time.Time
	if nullableCreatedAt != nil {
		createdAt = *nullableCreatedAt
	}
	var updatedAt time.Time
	if nullableUpdatedAt != nil {
		updatedAt = *nullableUpdatedAt
	}
	projectContainsUserObject, validationError := projectContainsUser.NewBuilder().WithUserID(userId).
		WithUserEmail(userEmail).WithFirstName(userFirstName).WithProjectID(projectId).
		WithLastName(userLastName).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).WithDeletedAt(deletedAt).Build()
	if validationError != nil {
		return nil, logger.LogCustomError(validationError)
	}
	return projectContainsUserObject, nil
}
