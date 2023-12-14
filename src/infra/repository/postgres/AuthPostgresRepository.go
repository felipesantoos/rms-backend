package postgres

import (
	"github.com/google/uuid"
	"rms-backend/src/core/domain/credentials"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/user"
	"rms-backend/src/core/interfaces/repository"
	"rms-backend/src/core/logger"
	"rms-backend/src/core/messages"
	"rms-backend/src/infra/repository/postgres/database"
	"rms-backend/src/infra/repository/postgres/queryObject"
)

type authPostgresRepository struct{}

func NewAuthPostgresRepository() repository.IAuthRepository {
	return &authPostgresRepository{}
}

func (instance authPostgresRepository) Login(credentials credentials.Credentials) (user.User, errors.Error) {
	row, err := database.Queryx(database.User().Query().ByEmail(), credentials.Email())
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	for !row.Next() {
		return nil, logger.LogCustomError(errors.NewFromString(messages.TheProvidedCredentialsAreInvalid))
	}
	serializedUser := map[string]interface{}{}
	row.MapScan(serializedUser)
	userObject, err := queryObject.NewUserFromMapRows(serializedUser)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	return userObject, nil
}

func (instance authPostgresRepository) SignUp(_user user.User) (*uuid.UUID, errors.Error) {
	transaction, err := database.BeginTransaction()
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	defer transaction.CloseConn()
	id, err := txQueryRowReturningID(
		transaction,
		database.User().Command().Create(),
		_user.Email(),
		_user.FirstName(),
		_user.LastName(),
		_user.Password(),
		_user.Salt(),
	)
	if err != nil {
		return nil, logger.LogCustomError(err)
	}
	userID, conversionError := uuid.Parse(id)
	if conversionError != nil {
		logger.LogNativeError(conversionError)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	err = transaction.Commit()
	if err != nil {
		logger.LogCustomError(err)
		return nil, logger.LogCustomError(errors.NewUnexpected())
	}
	return &userID, nil
}
