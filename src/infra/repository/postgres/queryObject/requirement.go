package queryObject

import (
	"fmt"
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/requirement"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/utils"
	"rms-backend/src/infra/repository/postgres/database"
	"strconv"
	"time"
)

func NewRequirementFromMapRows(data map[string]interface{}) (requirement.Requirement, errors.Error) {
	var err error
	id, err := uuid.Parse(string(data[database.RequirementID].([]uint8)))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	code, err := strconv.Atoi(fmt.Sprint(data[database.RequirementCode]))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	var title = fmt.Sprint(data[database.RequirementTitle])
	var description = fmt.Sprint(data[database.RequirementDescription])
	var userStory = fmt.Sprint(data[database.RequirementUserStory])
	typeID, err := uuid.Parse(string(data[database.RequirementTypeID].([]uint8)))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	originID, err := uuid.Parse(string(data[database.RequirementOriginID].([]uint8)))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	priorityID, err := uuid.Parse(string(data[database.RequirementPriorityID].([]uint8)))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	projectID, err := uuid.Parse(string(data[database.RequirementProjectID].([]uint8)))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	nullableCreatedAt := utils.GetNullableValue[time.Time](data[database.RequirementCreatedAt])
	nullableUpdatedAt := utils.GetNullableValue[time.Time](data[database.RequirementUpdatedAt])
	var createdAt time.Time
	if nullableCreatedAt != nil {
		createdAt = *nullableCreatedAt
	}
	var updatedAt time.Time
	if nullableUpdatedAt != nil {
		updatedAt = *nullableUpdatedAt
	}
	requirementObject, validationError := requirement.NewBuilder().WithID(id).WithCode(code).WithTitle(title).
		WithDescription(description).WithUserStory(userStory).WithTypeID(typeID).WithOriginID(originID).
		WithPriorityID(priorityID).WithProjectID(projectID).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).Build()
	if validationError != nil {
		return nil, logger.LogCustomError(validationError)
	}
	return requirementObject, nil
}
