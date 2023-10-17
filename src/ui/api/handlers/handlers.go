package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"rms-backend/src/core/domain/errors"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wallrony/go-validator/validator"
)

var unprocessableEntityError = &echo.HTTPError{
	Code: http.StatusUnprocessableEntity,
}
var unsupportedMediaTypeError = &echo.HTTPError{
	Message: "Unsupported Media Type",
	Code:    http.StatusUnsupportedMediaType,
}
var badRequestError = &echo.HTTPError{
	Code: http.StatusBadRequest,
}
var internalServerError = &echo.HTTPError{
	Code:    http.StatusInternalServerError,
	Message: "Ocorreu um erro inesperado. Por favor, contate o suporte.",
}
var unauthorizedError = &echo.HTTPError{
	Code: http.StatusUnauthorized,
}

func badRequestErrorWithMessage(message string) *echo.HTTPError {
	err := badRequestError
	err.Message = message
	return err
}

func unprocessableEntityErrorWithMessage(message string) *echo.HTTPError {
	err := unprocessableEntityError
	err.Message = message
	return err
}

func unsupportedMediaTypeErrorWithMessage(message string) *echo.HTTPError {
	err := unsupportedMediaTypeError
	err.Message = message
	return err
}

func responseFromError(err errors.Error) error {
	var e *echo.HTTPError = badRequestError
	if err.CausedInternally() {
		e = internalServerError
	} else if err.CausedByValidation() {
		e = unprocessableEntityError
	}
	e.Message = strings.Join(err.Messages(), ";")
	return e
}

func responseFromValidationError(valErr validator.ValidationError) error {
	var e *echo.HTTPError = badRequestError
	var err = errors.NewValidation(valErr.Messages())
	if err.CausedInternally() {
		e = internalServerError
	} else if err.CausedByValidation() {
		e = unprocessableEntityError
	}
	e.Message = strings.Join(err.Messages(), ";")
	return e
}

func getUUIDParamFromRequestPath(ctx echo.Context, paramName string) (*uuid.UUID, errors.Error) {
	strUUID := ctx.Param(paramName)
	if strUUID == "" {
		return nil, errors.NewFromString(fmt.Sprintf("you must provide a valid %s", paramName))
	} else if uuid, err := uuid.Parse(strUUID); err != nil {
		return nil, errors.NewFromString(fmt.Sprintf("the provided %s is invalid", paramName))
	} else {
		return &uuid, nil
	}
}

func getBoolQueryParamValue(ctx echo.Context, paramName string) (*bool, errors.Error) {
	value := ctx.QueryParam(paramName)
	parsedValue, err := strconv.ParseBool(value)
	if err != nil {
		return nil, errors.NewFromString(fmt.Sprintf("the provided %s is invalid", paramName))
	}

	return &parsedValue, nil
}

func getUUIDQueryParamValue(ctx echo.Context, paramName string) (*uuid.UUID, errors.Error) {
	value := ctx.QueryParam(paramName)
	parsedValue, err := uuid.Parse(value)
	if err != nil {
		return nil, errors.NewFromString(fmt.Sprintf("the provided %s is invalid", paramName))
	}

	return &parsedValue, nil
}

func getBoolFormDataValue(ctx echo.Context, fieldName string) (*bool, errors.Error) {
	value := ctx.FormValue(fieldName)
	parsedValue, err := strconv.ParseBool(value)
	if err != nil {
		return nil, errors.NewFromString(fmt.Sprintf("the provided %s is invalid", fieldName))
	}

	return &parsedValue, nil
}

func getIntFormDataValue(ctx echo.Context, fieldName string) (*int, errors.Error) {
	value := ctx.FormValue(fieldName)
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return nil, errors.NewFromString(fmt.Sprintf("the provided %s is invalid", fieldName))
	}

	return &parsedValue, nil
}

func getTimeFormDataValue(ctx echo.Context, fieldName string) (*time.Time, errors.Error) {
	value := ctx.FormValue(fieldName)
	parsedValue, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return nil, errors.NewFromString(fmt.Sprintf("the provided %s is invalid", fieldName))
	}

	return &parsedValue, nil
}

func getUUIDFormDataValue(ctx echo.Context, fieldName string) (*uuid.UUID, errors.Error) {
	value := ctx.FormValue(fieldName)
	parsedValue, err := uuid.Parse(value)
	if err != nil {
		return nil, errors.NewFromString(fmt.Sprintf("the provided %s is invalid", fieldName))
	}

	return &parsedValue, nil
}

func stringListToUUIDList(list []string) ([]uuid.UUID, errors.Error) {
	ids := []uuid.UUID{}
	for _, item := range list {
		id, err := uuid.Parse(item)
		if err != nil {
			return nil, errors.NewFromString(fmt.Sprintf("the provided string %s cannot be converted to uuid", item))
		}
		ids = append(ids, id)
	}
	return ids, nil
}
