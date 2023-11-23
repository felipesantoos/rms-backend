package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"rms-backend/src/core/domain/errors"
	primary "rms-backend/src/core/interfaces/primary"
	"rms-backend/src/core/services/filters"
	"rms-backend/src/ui/api/handlers/dto/request"
	"rms-backend/src/ui/api/handlers/dto/response"
	"rms-backend/src/ui/api/handlers/utils/checkers"
	"rms-backend/src/ui/api/handlers/utils/params"
)

type RequirementHandlers interface {
	Create(echo.Context) error
	List(echo.Context) error
	Get(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

type requirementHandlers struct {
	requirementServices primary.IRequirementServices
}

func NewRequirementHandler(requirementServices primary.IRequirementServices) RequirementHandlers {
	return &requirementHandlers{requirementServices: requirementServices}
}

// Create
// @ID Requirement.Create
// @Summary Cadastrar requisito
// @Description Rota que permite o cadastro de um requisito de um projeto.
// @Tags Requisitos
// @Accept json
// @Param json body request.Requirement true "JSON com todos os dados necessários para se cadastrar um requisito."
// @Produce json
// @Success 201 {object} response.ID "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /requirements [post]
func (this *requirementHandlers) Create(ctx echo.Context) error {
	var requirementDTO request.Requirement
	bindError := ctx.Bind(&requirementDTO)
	if bindError != nil {
		return responseFromError(errors.New(bindError))
	}
	requirementObject, validationError := requirementDTO.ToDomain()
	if validationError != nil {
		return responseFromError(validationError)
	}
	id, err := this.requirementServices.Create(requirementObject)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusCreated, response.IDBuilder().FromUUID(*id))
}

// List
// @ID Requirement.List
// @Summary Listar requisitos de um projeto
// @Description Rota que permite a listagem dos requisitos de um projeto.
// @Tags Requisitos
// @Param projectID query string false "ID do projeto."
// @Produce json
// @Success 200 {array} response.Requirement "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /requirements [get]
func (this *requirementHandlers) List(ctx echo.Context) error {
	var projectID *uuid.UUID
	if !checkers.IsEmpty(ctx.QueryParam(params.ProjectID)) {
		aux, err := getUUIDQueryParamValue(ctx, params.ProjectID)
		if err != nil {
			return responseFromError(err)
		}
		projectID = aux
	}
	requirementFilters := filters.RequirementFilters{ProjectID: projectID}
	requirements, err := this.requirementServices.List(requirementFilters)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.RequirementBuilder().BuildFromDomainList(requirements))
}

// Get
// @ID Requirement.Get
// @Summary Buscar detalhes de um projeto pelo ID
// @Description Rota que permite a buscar dos detalhes de um requisito pelo ID.
// @Tags Requisitos
// @Param id path string true "ID do requisito."
// @Produce json
// @Success 200 {object} response.Requirement "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /requirements/{id} [get]
func (this *requirementHandlers) Get(ctx echo.Context) error {
	id, err := getUUIDParamFromRequestPath(ctx, params.ID)
	requirement, err := this.requirementServices.Get(*id)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.RequirementBuilder().BuildFromDomain(requirement))
}

// Update
// @ID Requirement.Update
// @Summary Atualizar requisito
// @Description Rota que permite a atualização de um requisito de um projeto.
// @Tags Requisitos
// @Accept json
// @Param id path string true "ID do requisito."
// @Param json body request.Requirement true "JSON com todos os dados necessários para se atualizar um projeto."
// @Produce json
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /requirements/{id} [put]
func (this *requirementHandlers) Update(ctx echo.Context) error {
	id, err := getUUIDParamFromRequestPath(ctx, params.ID)
	var requirementDTO request.Requirement
	bindError := ctx.Bind(&requirementDTO)
	if bindError != nil {
		return responseFromError(errors.New(bindError))
	}
	requirementObject, validationError := requirementDTO.ToDomain()
	if validationError != nil {
		return responseFromError(validationError)
	}
	validationError = requirementObject.SetID(*id)
	if validationError != nil {
		return responseFromError(validationError)
	}
	err = this.requirementServices.Update(requirementObject)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.NoContent(http.StatusNoContent)
}

// Delete
// @ID Requirement.Delete
// @Summary Deletar requisito de um projeto
// @Description Rota que permite a deleção de um requisito de um projeto.
// @Tags Requisitos
// @Accept json
// @Param id path string true "ID do requisito."
// @Produce json
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /requirements/{id} [delete]
func (this *requirementHandlers) Delete(ctx echo.Context) error {
	id, err := getUUIDParamFromRequestPath(ctx, params.ID)
	err = this.requirementServices.Delete(*id)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.NoContent(http.StatusNoContent)
}
