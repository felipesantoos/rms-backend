package infra

import (
	"errors"
	"fmt"
)

type Error interface {
	Err() string
	IsInternal() bool
	Native() error
}

type sourceErr struct {
	err        error
	isInternal bool
}

func NewSourceErr(err error) Error {
	return &sourceErr{err: err}
}

func NewSourceErrFromStr(message string) Error {
	err := errors.New(message)
	return &sourceErr{err: err}
}

func NewInternalSourceErr(err error) Error {
	return &sourceErr{err, true}
}

func NewUnexpectedSourceErr() Error {
	err := errors.New(fmt.Sprintf("An unexpected error occurred. Please contact the support."))
	return &sourceErr{err, true}
}

func (e *sourceErr) Err() string {
	return e.err.Error()
}

func (e *sourceErr) IsInternal() bool {
	return e.isInternal
}

func (e *sourceErr) Native() error {
	return e.err
}
