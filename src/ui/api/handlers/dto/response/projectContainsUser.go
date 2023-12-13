package response

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/projectContainsUser"
	"time"
)

type ProjectContainsUser struct {
	UserID        uuid.UUID  `json:"user_id"`
	UserEmail     string     `json:"user_email"`
	UserFirstName string     `json:"user_first_name"`
	UserLastName  string     `json:"user_last_name"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

type projectContainsUserBuilder struct{}

func ProjectContainsUserBuilder() *projectContainsUserBuilder {
	return &projectContainsUserBuilder{}
}

func (*projectContainsUserBuilder) BuildFromDomain(data projectContainsUser.ProjectContainsUser) ProjectContainsUser {
	return ProjectContainsUser{
		UserID:        data.UserID(),
		UserEmail:     data.UserEmail(),
		UserFirstName: data.UserFirstName(),
		UserLastName:  data.UserLastName(),
		CreatedAt:     data.CreatedAt(),
		UpdatedAt:     data.UpdatedAt(),
		DeletedAt:     data.DeletedAt(),
	}
}

func (builder *projectContainsUserBuilder) BuildFromDomainList(list []projectContainsUser.ProjectContainsUser) []ProjectContainsUser {
	result := []ProjectContainsUser{}
	for _, item := range list {
		result = append(result, builder.BuildFromDomain(item))
	}
	return result
}
