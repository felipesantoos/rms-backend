package queryObject

import (
	"fmt"
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/type"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/utils"
	"rms-backend/src/infra/repository/postgres/database"
	"time"
)

func NewTypeFromMapRows(data map[string]interface{}) (_type.Type, errors.Error) {
	var err error
	id, err := uuid.Parse(string(data[database.TypeID].([]uint8)))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	var name = fmt.Sprint(data[database.TypeName])
	var description = fmt.Sprint(data[database.TypeDescription])
	nullableCreatedAt := utils.GetNullableValue[time.Time](data[database.TypeCreatedAt])
	nullableUpdatedAt := utils.GetNullableValue[time.Time](data[database.TypeUpdatedAt])
	var createdAt time.Time
	if nullableCreatedAt != nil {
		createdAt = *nullableCreatedAt
	}
	var updatedAt time.Time
	if nullableUpdatedAt != nil {
		updatedAt = *nullableUpdatedAt
	}
	typeObject, validationError := _type.NewBuilder().WithID(id).WithName(name).WithDescription(description).
		WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).Build()
	if validationError != nil {
		return nil, logger.LogCustomError(validationError)
	}
	return typeObject, nil
}
