package request

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/project"
	"rms-backend/src/core/logger"
)

type Project struct {
	Name        string `json:"name" example:"Project A"`
	Alias       string `json:"alias" example:"PA"`
	Description string `json:"description" example:"Description of the Project A."`
	IsActive    bool   `json:"is_active" example:"true"`
}

func (this *Project) ToDomain() (project.Project, errors.Error) {
	projectObject, validationError := project.NewBuilder().WithName(this.Name).WithAlias(this.Alias).
		WithDescription(this.Description).WithIsActive(this.IsActive).Build()
	if validationError != nil {
		return nil, logger.LogCustomError(validationError)
	}
	return projectObject, nil
}
