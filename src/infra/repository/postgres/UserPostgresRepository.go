package postgres

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/user"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/messages"
	"rms-backend/src/infra/repository/postgres/database"
	"rms-backend/src/infra/repository/postgres/queryObject"
)

type userPostgresRepository struct{}

func NewUserPostgresRepository() repository.IUserRepository {
	return &userPostgresRepository{}
}

func (this *userPostgresRepository) GetByEmail(email string) (user.User, errors.Error) {
	row, err := database.Queryx(database.User().Query().ByEmail(), email)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	for !row.Next() {
		return nil, logger.LogCustomError(errors.NewFromString(messages.UserNotFoundErrorMessage))
	}
	serializedUser := map[string]interface{}{}
	row.MapScan(serializedUser)
	userObject, err := queryObject.NewUserFromMapRows(serializedUser)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return userObject, nil
}
