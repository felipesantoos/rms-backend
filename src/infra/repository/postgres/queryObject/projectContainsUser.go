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

	var createdAt = utils.GetNullableValue[time.Time](data[database.ProjectContainsUserCreatedAt])
	var updatedAt = utils.GetNullableValue[time.Time](data[database.ProjectContainsUserUpdatedAt])
	var deletedAt = utils.GetNullableValue[time.Time](data[database.ProjectContainsUserDeletedAt])

	projectContainsUserObject, validationError := projectContainsUser.NewBuilder().WithUserID(userId).
		WithUserEmail(userEmail).WithFirstName(userFirstName).
		WithLastName(userLastName).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).WithDeletedAt(deletedAt).Build()
	if validationError != nil {
		return nil, logger.LogCustomError(validationError)
	}
	return projectContainsUserObject, nil
}
