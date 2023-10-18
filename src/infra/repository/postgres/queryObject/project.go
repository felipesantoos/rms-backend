package queryObject

import (
	"fmt"
	"github.com/google/uuid"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/knowledgeArea"
	"rms-backend/src/core/logger"
)

func NewKnowledgeAreaFromMapRows(data map[string]interface{}) (knowledgeArea.KnowledgeArea, errors.Error) {
	var err error
	var id uuid.UUID
	var name = fmt.Sprint(data[""])
	id, err = uuid.Parse(string(data[""].([]uint8)))
	if err != nil {
		logger.LogNativeError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	_knowledgeArea, validationError := knowledgeArea.NewBuilder().WithID(id).WithName(name).Build()
	if err != nil {
		return nil, logger.LogCustomError(validationError)
	}
	return _knowledgeArea, nil
}
