package credentials

import (
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/messages"
)

type builder struct {
	fields        []string
	errorMessages []string
	credentials   *credentials
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		credentials:   &credentials{},
	}
}

func (this *builder) WithEmail(email string) *builder {
	this.credentials.email = email
	return this
}

func (this *builder) WithPassword(password string) *builder {
	this.credentials.password = password
	return this
}

func (this *builder) Build() (*credentials, errors.Error) {
	if len(this.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(this.errorMessages, map[string]interface{}{
			messages.Fields: this.fields})
	}
	return this.credentials, nil
}
