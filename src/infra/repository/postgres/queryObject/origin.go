package queryObject

import (
	"fmt"
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/origin"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/utils"
	"rms-backend/src/infra/repository/postgres/database"
	"time"
)

func NewOriginFromMapRows(data map[string]interface{}) (origin.Origin, errors.Error) {
	var err error
	id, err := uuid.Parse(string(data[database.OriginID].([]uint8)))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	var name = fmt.Sprint(data[database.OriginName])
	var description = fmt.Sprint(data[database.OriginDescription])
	nullableCreatedAt := utils.GetNullableValue[time.Time](data[database.OriginCreatedAt])
	nullableUpdatedAt := utils.GetNullableValue[time.Time](data[database.OriginUpdatedAt])
	var createdAt time.Time
	if nullableCreatedAt != nil {
		createdAt = *nullableCreatedAt
	}
	var updatedAt time.Time
	if nullableUpdatedAt != nil {
		updatedAt = *nullableUpdatedAt
	}
	originObject, validationError := origin.NewBuilder().WithID(id).WithName(name).WithDescription(description).
		WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).Build()
	if validationError != nil {
		return nil, logger.LogCustomError(validationError)
	}
	return originObject, nil
}
