package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rms-backend/src/core/domain/credentials"
	primary "rms-backend/src/core/interfaces/primary"
	"rms-backend/src/ui/api/handlers/dto/request"
	"rms-backend/src/ui/api/handlers/dto/response"
)

type AuthHandlers interface {
	Login(echo.Context) error
}

type authHandlers struct {
	authServices primary.IAuthServices
}

func NewAuthHandler(authServices primary.IAuthServices) AuthHandlers {
	return &authHandlers{authServices: authServices}
}

// Login
// @ID Login
// @Summary Fazer login no sistema
// @Tags Rotas de autenticação
// @Description Rota que permite que um usuário se autentique no sistema utilizando seu endereço de e-mail e senha.
// @Description | E-mail              | Senha     |
// @Description |---------------------|-----------|
// @Description | john@gmail.com      | Test1234! |
// @Description | jane@gmail.com      | Test1234! |
// @Description | bob@gmail.com       | Test1234! |
// @Description | williams@gmail.com  | Test1234! |
// @Description | brown@gmail.com     | Test1234! |
// @Accept json
// @Produce json
// @Param json body request.LoginDTO true "JSON com todos os dados necessários para que o login seja realizado."
// @Success 201 {object} response.Authorization "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /auth/login [post]
func (instance authHandlers) Login(context echo.Context) error {
	var loginDTO request.LoginDTO
	bindError := context.Bind(&loginDTO)
	if bindError != nil {
		return badRequestErrorWithMessage("Não foi possível processar a solicitação")
	}
	_credentials, err := credentials.NewBuilder().WithEmail(loginDTO.Email).WithPassword(loginDTO.Password).Build()
	if err != nil {
		return responseFromError(err)
	}
	_authorization, err := instance.authServices.Login(_credentials)
	if err != nil {
		return responseFromError(err)
	}
	return context.JSON(http.StatusCreated, response.AuthorizationBuilder().BuildFromDomain(_authorization))
}
