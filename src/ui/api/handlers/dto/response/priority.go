package response

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/priority"
	"time"
)

type Priority struct {
	ID          uuid.UUID `json:"id"`
	Level       string    `json:"level"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type priorityBuilder struct{}

func PriorityBuilder() *priorityBuilder {
	return &priorityBuilder{}
}

func (*priorityBuilder) BuildFromDomain(data priority.Priority) Priority {
	return Priority{
		ID:          data.ID(),
		Level:       data.Level(),
		Description: data.Description(),
		CreatedAt:   data.CreatedAt(),
		UpdatedAt:   data.UpdatedAt(),
	}
}

func (builder *priorityBuilder) BuildFromDomainList(list []priority.Priority) []Priority {
	result := []Priority{}
	for _, item := range list {
		result = append(result, builder.BuildFromDomain(item))
	}
	return result
}
