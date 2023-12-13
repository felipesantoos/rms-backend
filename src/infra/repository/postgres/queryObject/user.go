package queryObject

import (
	"fmt"
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/user"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/utils"
	"rms-backend/src/infra/repository/postgres/database"
	"strconv"
	"time"
)

func NewUserFromMapRows(data map[string]interface{}) (user.User, errors.Error) {
	var err error
	id, err := uuid.Parse(string(data[database.UserID].([]uint8)))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	var email = fmt.Sprint(data[database.UserEmail])
	var firstName = fmt.Sprint(data[database.UserFirstName])
	var lastName = fmt.Sprint(data[database.UserLastName])
	isActive, err := strconv.ParseBool(fmt.Sprint(data[database.UserIsActive]))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	nullableCreatedAt := utils.GetNullableValue[time.Time](data[database.UserCreatedAt])
	nullableUpdatedAt := utils.GetNullableValue[time.Time](data[database.UserUpdatedAt])
	var deletedAt = utils.GetNullableValue[time.Time](data[database.UserDeletedAt])
	var createdAt time.Time
	if nullableCreatedAt != nil {
		createdAt = *nullableCreatedAt
	}
	var updatedAt time.Time
	if nullableUpdatedAt != nil {
		updatedAt = *nullableUpdatedAt
	}
	userObject, validationError := user.NewBuilder().WithID(id).WithEmail(email).WithFirstName(firstName).
		WithLastName(lastName).WithIsActive(isActive).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).
		WithDeletedAt(deletedAt).Build()
	if validationError != nil {
		return nil, logger.LogCustomError(validationError)
	}
	return userObject, nil
}
