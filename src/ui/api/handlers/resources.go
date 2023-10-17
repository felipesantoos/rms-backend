package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rms-backend/src/core/interfaces/usecases"
	"rms-backend/src/ui/api/handlers/dto/response"
)

type ResourcesHandler interface {
	ListKnowledgeAreas(echo.Context) error
}

type resourcesHandler struct {
	services usecases.ResourcesUseCase
}

func NewResourcesHandler(services usecases.ResourcesUseCase) ResourcesHandler {
	return &resourcesHandler{services}
}

// ListKnowledgeAreas
// List Knowledge Areas
// @ID Resources.ListKnowledgeAreas
// @Summary Listar todas as áreas de conhecimento cadastradas no banco de dados.
// @Description Rota que lista todas as áreas de conhecimento cadastradas no banco de dados.
// @Tags Rotas de recursos
// @Produce json
// @Success 200 {array} response.KnowledgeArea "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /res/knowledge-areas [get]
func (this *resourcesHandler) ListKnowledgeAreas(ctx echo.Context) error {
	result, err := this.services.ListKnowledgeAreas()
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.KnowledgeAreaBuilder().BuildFromDomainList(result))
}
