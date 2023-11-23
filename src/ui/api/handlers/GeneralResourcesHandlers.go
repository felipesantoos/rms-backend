package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	primary "rms-backend/src/core/interfaces/primary"
	"rms-backend/src/ui/api/handlers/dto/response"
)

type GeneralResourcesHandlers interface {
	List(echo.Context) error
}

type generalResourcesHandlers struct {
	originServices   primary.IOriginServices
	priorityServices primary.IPriorityServices
	typeServices     primary.ITypeServices
}

func NewGeneralResourcesHandlers(originServices primary.IOriginServices, priorityServices primary.IPriorityServices,
	typeServices primary.ITypeServices) GeneralResourcesHandlers {
	return &generalResourcesHandlers{
		originServices:   originServices,
		priorityServices: priorityServices,
		typeServices:     typeServices,
	}
}

// List
// @ID GeneralResources.List
// @Summary Listar todos os recursos
// @Description Rota que permite a listagem de todos os recursos.
// @Tags Recursos
// @Produce json
// @Success 200 {array} response.GeneralResources "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /resources [get]
func (this *generalResourcesHandlers) List(ctx echo.Context) error {
	var generalResources response.GeneralResources
	origins, err := this.originServices.List()
	if err != nil {
		return responseFromError(err)
	}
	generalResources.Origins = response.OriginBuilder().BuildFromDomainList(origins)
	priorities, err := this.priorityServices.List()
	if err != nil {
		return responseFromError(err)
	}
	generalResources.Priorities = response.PriorityBuilder().BuildFromDomainList(priorities)
	types, err := this.typeServices.List()
	if err != nil {
		return responseFromError(err)
	}
	generalResources.Types = response.TypeBuilder().BuildFromDomainList(types)
	return ctx.JSON(http.StatusOK, generalResources)
}
