package response

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/origin"
	"time"
)

type Origin struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type originBuilder struct{}

func OriginBuilder() *originBuilder {
	return &originBuilder{}
}

func (*originBuilder) BuildFromDomain(data origin.Origin) Origin {
	return Origin{
		ID:          data.ID(),
		Name:        data.Name(),
		Description: data.Description(),
		CreatedAt:   data.CreatedAt(),
		UpdatedAt:   data.UpdatedAt(),
	}
}

func (builder *originBuilder) BuildFromDomainList(list []origin.Origin) []Origin {
	result := []Origin{}
	for _, item := range list {
		result = append(result, builder.BuildFromDomain(item))
	}
	return result
}
