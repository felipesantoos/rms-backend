package token

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"rms-backend/src/core/domain/authorization"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/messages"
	"rms-backend/src/core/utils"
)

func getAuthClaims(authHeader string) (*authorization.Claims, errors.Error) {
	_, token := utils.ExtractToken(authHeader)
	authClaims, err := utils.ExtractTokenClaims(token)
	if err != nil {
		return nil, errors.NewValidationFromString(messages.InvalidTokenErrorMessage)
	}
	return authClaims, nil
}

func GetAccountIDFromAuthorization(ctx echo.Context) (*uuid.UUID, errors.Error) {
	claims, err := getAuthClaims(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		return nil, err
	}
	if accountID, parseError := uuid.Parse(claims.AccountID); parseError != nil {
		return nil, errors.NewValidationFromString(messages.InvalidIDFromTokenErrorMessage)
	} else {
		return &accountID, nil
	}
}
