package response

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/project"
	"time"
)

type Project struct {
	ID                 uuid.UUID  `json:"id"`
	Name               string     `json:"name"`
	Alias              string     `json:"alias"`
	Description        string     `json:"description"`
	IsActive           bool       `json:"is_active"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
	CreatedByUserEmail string     `json:"created_by_user_email"`
}

type projectBuilder struct{}

func ProjectBuilder() *projectBuilder {
	return &projectBuilder{}
}

func (*projectBuilder) BuildFromDomain(data project.Project) Project {
	return Project{
		ID:                 data.ID(),
		Name:               data.Name(),
		Alias:              data.Alias(),
		Description:        data.Description(),
		IsActive:           data.IsActive(),
		CreatedAt:          data.CreatedAt(),
		UpdatedAt:          data.UpdatedAt(),
		DeletedAt:          data.DeletedAt(),
		CreatedByUserEmail: data.CreatedByUserEmail(),
	}
}

func (builder *projectBuilder) BuildFromDomainList(list []project.Project) []Project {
	result := []Project{}
	for _, item := range list {
		result = append(result, builder.BuildFromDomain(item))
	}
	return result
}
