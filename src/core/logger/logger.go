package logger

import (
	"log"
	"rms-backend/src/core/domain/errors"
)

func LogCustomError(err errors.Error) errors.Error {
	log.Println(err)
	return err
}

func LogNativeError(err error) {
	log.Println(err)
}
