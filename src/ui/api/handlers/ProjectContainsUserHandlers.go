package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rms-backend/src/core/domain/errors"
	"rms-backend/src/core/interfaces/primary"
	"rms-backend/src/ui/api/handlers/dto/request"
	"rms-backend/src/ui/api/handlers/dto/response"
	"rms-backend/src/ui/api/handlers/utils/params"
)

type ProjectContainsUserHandlers interface {
	Create(echo.Context) error
	List(echo.Context) error
	Delete(echo.Context) error
}

type projectContainsUserHandlers struct {
	projectContainsUserServices primary.IProjectContainsUserServices
	userServices                primary.IUserServices
}

func NewProjectContainsUserHandlers(projectContainsUserServices primary.IProjectContainsUserServices, userServices primary.IUserServices) ProjectContainsUserHandlers {
	return &projectContainsUserHandlers{
		projectContainsUserServices: projectContainsUserServices,
		userServices:                userServices,
	}
}

// Create
// @ID ProjectContainsUser.Create
// @Summary Adicionar colaborador
// @Description Rota que permite a adição de um colaborador ao projeto.
// @Tags Colaboradores
// @Accept json
// @Param projectID path string true "ID do projeto."
// @Param json body request.ProjectContainsUser true "JSON com todos os dados necessários para adicionar um colaborador."
// @Produce json
// @Success 201 {object} response.ID "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /collaborators/{projectID} [post]
func (this *projectContainsUserHandlers) Create(ctx echo.Context) error {
	projectID, err := getUUIDParamFromRequestPath(ctx, params.ProjectID)
	if err != nil {
		return responseFromError(err)
	}

	var projectContainsUserDTO request.ProjectContainsUser
	bindError := ctx.Bind(&projectContainsUserDTO)
	if bindError != nil {
		return responseFromError(errors.New(bindError))
	}

	userInstance, err := this.userServices.GetByEmail(projectContainsUserDTO.UserEmail)
	if err != nil {
		return responseFromError(err)
	}

	projectContainsUserObject, validationError := projectContainsUserDTO.ToDomain(userInstance)
	if validationError != nil {
		return responseFromError(validationError)
	}
	err = this.projectContainsUserServices.Create(*projectID, projectContainsUserObject)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.NoContent(http.StatusCreated)
}

// List
// @ID ProjectContainsUser.List
// @Summary Listar colaboradores
// @Description Rota que permite a listagem dos colaboradores.
// @Tags Colaboradores
// @Param projectID path string true "ID do projeto."
// @Produce json
// @Success 200 {array} response.ProjectContainsUser "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /collaborators/{projectID} [get]
func (this *projectContainsUserHandlers) List(ctx echo.Context) error {
	projectID, err := getUUIDParamFromRequestPath(ctx, params.ProjectID)
	if err != nil {
		return responseFromError(err)
	}

	projectContainsUsers, err := this.projectContainsUserServices.List(*projectID)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.ProjectContainsUserBuilder().BuildFromDomainList(projectContainsUsers))
}

// Delete
// @ID ProjectContainsUser.Delete
// @Summary Deletar colaborador
// @Description Rota que permite a deleção de um colaborador.
// @Tags Colaboradores
// @Accept json
// @Param projectID path string true "ID do projeto."
// @Param userID path string true "ID do usuário colaborador."
// @Produce json
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /collaborators/{projectID}/{userID} [delete]
func (this *projectContainsUserHandlers) Delete(ctx echo.Context) error {
	projectID, err := getUUIDParamFromRequestPath(ctx, params.ProjectID)
	if err != nil {
		return responseFromError(err)
	}

	userID, err := getUUIDParamFromRequestPath(ctx, params.UserID)
	if err != nil {
		return responseFromError(err)
	}

	err = this.projectContainsUserServices.Delete(*projectID, *userID)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.NoContent(http.StatusNoContent)
}
