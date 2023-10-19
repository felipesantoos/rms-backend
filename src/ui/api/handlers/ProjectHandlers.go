package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rms-backend/src/core/domain/errors"
	primary "rms-backend/src/core/interfaces/primary"
	"rms-backend/src/ui/api/handlers/dto/request"
	"rms-backend/src/ui/api/handlers/dto/response"
	"rms-backend/src/ui/api/handlers/utils/params"
)

type ProjectHandlers interface {
	Create(echo.Context) error
	List(echo.Context) error
	Get(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

type projectHandlers struct {
	projectServices primary.IProjectServices
}

func NewResourcesHandler(projectServices primary.IProjectServices) ProjectHandlers {
	return &projectHandlers{projectServices: projectServices}
}

// Create
// @ID Project.Create
// @Summary Cadastrar projeto
// @Description Rota que permite o cadastro de um projeto.
// @Tags Projetos
// @Accept json
// @Param json body request.Project true "JSON com todos os dados necessários para se cadastrar um projeto."
// @Produce json
// @Success 201 {object} response.ID "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /projects [post]
func (this *projectHandlers) Create(ctx echo.Context) error {
	var projectDTO request.Project
	bindError := ctx.Bind(&projectDTO)
	if bindError != nil {
		return responseFromError(errors.New(bindError))
	}
	projectObject, validationError := projectDTO.ToDomain()
	if validationError != nil {
		return responseFromError(validationError)
	}
	id, err := this.projectServices.Create(projectObject)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusCreated, response.IDBuilder().FromUUID(*id))
}

// List
// @ID Project.List
// @Summary Listar projetos
// @Description Rota que permite a listagem dos projetos.
// @Tags Projetos
// @Produce json
// @Success 200 {array} response.Project "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /projects [get]
func (this *projectHandlers) List(ctx echo.Context) error {
	projects, err := this.projectServices.List()
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.ProjectBuilder().BuildFromDomainList(projects))
}

// Get
// @ID Project.Get
// @Summary Buscar detalhes de um projeto pelo ID
// @Description Rota que permite a buscar dos detalhes de um projeto pelo ID.
// @Tags Projetos
// @Param id path string true "ID do projeto."
// @Produce json
// @Success 200 {object} response.Project "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /projects/{id} [get]
func (this *projectHandlers) Get(ctx echo.Context) error {
	id, err := getUUIDParamFromRequestPath(ctx, params.ID)
	project, err := this.projectServices.Get(*id)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.ProjectBuilder().BuildFromDomain(project))
}

func (this *projectHandlers) Update(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (this *projectHandlers) Delete(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
