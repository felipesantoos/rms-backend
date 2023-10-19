package queryObject

import (
	"fmt"
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/priority"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/utils"
	"rms-backend/src/infra/repository/postgres/database"
	"time"
)

func NewPriorityFromMapRows(data map[string]interface{}) (priority.Priority, errors.Error) {
	var err error
	id, err := uuid.Parse(string(data[database.PriorityID].([]uint8)))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	var name = fmt.Sprint(data[database.PriorityLevel])
	var description = fmt.Sprint(data[database.PriorityDescription])
	nullableCreatedAt := utils.GetNullableValue[time.Time](data[database.PriorityCreatedAt])
	nullableUpdatedAt := utils.GetNullableValue[time.Time](data[database.PriorityUpdatedAt])
	var createdAt time.Time
	if nullableCreatedAt != nil {
		createdAt = *nullableCreatedAt
	}
	var updatedAt time.Time
	if nullableUpdatedAt != nil {
		updatedAt = *nullableUpdatedAt
	}
	priorityObject, validationError := priority.NewBuilder().WithID(id).WithLevel(name).WithDescription(description).
		WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).Build()
	if validationError != nil {
		return nil, logger.LogCustomError(validationError)
	}
	return priorityObject, nil
}
