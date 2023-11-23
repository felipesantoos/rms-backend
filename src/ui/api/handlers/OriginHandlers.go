package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	primary "rms-backend/src/core/interfaces/primary"
	"rms-backend/src/ui/api/handlers/dto/response"
)

type OriginHandlers interface {
	List(echo.Context) error
}

type originHandlers struct {
	originServices primary.IOriginServices
}

func NewOriginHandler(originServices primary.IOriginServices) OriginHandlers {
	return &originHandlers{originServices: originServices}
}

// List
// @ID Origin.List
// @Summary Listar origens
// @Description Rota que permite a listagem das origens.
// @Tags Recursos
// @Produce json
// @Success 200 {array} response.Origin "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /resources/origins [get]
func (this *originHandlers) List(ctx echo.Context) error {
	origins, err := this.originServices.List()
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.OriginBuilder().BuildFromDomainList(origins))
}
