package response

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/type"
	"time"
)

type Type struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type typeBuilder struct{}

func TypeBuilder() *typeBuilder {
	return &typeBuilder{}
}

func (*typeBuilder) BuildFromDomain(data _type.Type) Type {
	return Type{
		ID:          data.ID(),
		Name:        data.Name(),
		Description: data.Description(),
		CreatedAt:   data.CreatedAt(),
		UpdatedAt:   data.UpdatedAt(),
	}
}

func (builder *typeBuilder) BuildFromDomainList(list []_type.Type) []Type {
	result := []Type{}
	for _, item := range list {
		result = append(result, builder.BuildFromDomain(item))
	}
	return result
}
