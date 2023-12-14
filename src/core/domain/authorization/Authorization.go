package authorization

import (
	"os"
	"rms-backend/src/core"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/domain/user"
	"time"

	"github.com/golang-jwt/jwt"
)

var logger = core.Logger()

const (
	AccessTokenTimeout = time.Minute * 60
	AnonymousRoleCode  = "anonymous"
	Type               = "Bearer"
)

type Authorization interface {
	Token() string
}

type authorization struct {
	token string
}

func New() Authorization {
	return &authorization{}
}

func NewFromAccount(account user.User) (Authorization, errors.Error) {
	instance := &authorization{}
	if err := instance.GenerateToken(account); err != nil {
		return nil, err
	}
	return instance, nil
}

func (instance *authorization) Token() string {
	return instance.token
}

func (instance *authorization) GenerateToken(_user user.User) errors.Error {
	secret := os.Getenv("SERVER_SECRET")
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims(
		_user.ID().String(),
		Type,
		time.Now().Add(AccessTokenTimeout).Unix(),
	)).SignedString([]byte(secret))
	if err != nil {
		logger.Error().Msg(err.Error())
		return errors.NewInternal(err)
	}
	instance.token = accessToken
	return nil
}
