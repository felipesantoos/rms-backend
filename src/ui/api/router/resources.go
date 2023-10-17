package router

import (
	"github.com/labstack/echo/v4"
	"rms-backend/src/ui/api/dicontainer"
	"rms-backend/src/ui/api/handlers"
)

type resourcesRouter struct {
	handler handlers.ResourcesHandler
}

func NewResourcesRouter() Router {
	usecase := dicontainer.ResourcesUseCase()
	handler := handlers.NewResourcesHandler(usecase)
	return &resourcesRouter{handler}
}

func (this *resourcesRouter) Load(apiGroup *echo.Group) {
	router := apiGroup.Group("/res")
	router.GET("/knowledge-areas", this.handler.ListKnowledgeAreas)
}
