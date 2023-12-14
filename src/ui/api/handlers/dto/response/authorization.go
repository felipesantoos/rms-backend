package response

import (
	"rms-backend/src/core/domain/authorization"
)

type Authorization struct {
	Token string `json:"access_token"`
}

type authorizationBuilder struct{}

func AuthorizationBuilder() *authorizationBuilder {
	return &authorizationBuilder{}
}

func (*authorizationBuilder) BuildFromDomain(data authorization.Authorization) Authorization {
	return Authorization{
		Token: data.Token(),
	}
}
