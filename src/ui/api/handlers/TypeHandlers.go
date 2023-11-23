package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	primary "rms-backend/src/core/interfaces/primary"
	"rms-backend/src/ui/api/handlers/dto/response"
)

type TypeHandlers interface {
	List(echo.Context) error
}

type typeHandlers struct {
	typeServices primary.ITypeServices
}

func NewTypeHandler(typeServices primary.ITypeServices) TypeHandlers {
	return &typeHandlers{typeServices: typeServices}
}

// List
// @ID Type.List
// @Summary Listar tipos
// @Description Rota que permite a listagem dos tipos.
// @Tags Recursos
// @Produce json
// @Success 200 {array} response.Type "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /types [get]
func (this *typeHandlers) List(ctx echo.Context) error {
	types, err := this.typeServices.List()
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.TypeBuilder().BuildFromDomainList(types))
}
