package request

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/projectContainsUser"
	"rms-backend/src/core/domain/user"
	"rms-backend/src/core/logger"
)

type ProjectContainsUser struct {
	UserEmail string `json:"user_email" example:"collaborator@gmail.com"`
}

func (this *ProjectContainsUser) ToDomain(userInstance user.User) (projectContainsUser.ProjectContainsUser, errors.Error) {
	projectContainsUserObject, validationError := projectContainsUser.NewBuilder().WithUserID(userInstance.ID()).Build()
	if validationError != nil {
		return nil, logger.LogCustomError(validationError)
	}
	return projectContainsUserObject, nil
}
