package router

import (
	"github.com/labstack/echo/v4"
	"rms-backend/src/ui/api/dicontainer"
	"rms-backend/src/ui/api/handlers"
)

type resourceRouter struct {
	generalResourcesHandler handlers.GeneralResourcesHandlers
}

func NewResourceRouter() Router {
	originServices := dicontainer.OriginServices()
	priorityServices := dicontainer.PriorityServices()
	typeServices := dicontainer.TypeServices()
	generalResourcesHandlers := handlers.NewGeneralResourcesHandlers(originServices, priorityServices, typeServices)
	return &resourceRouter{generalResourcesHandler: generalResourcesHandlers}
}

func (this *resourceRouter) Load(apiGroup *echo.Group) {
	resourceGroup := apiGroup.Group("/resources")
	resourceGroup.GET("", this.generalResourcesHandler.List)
	NewOriginRouter().Load(resourceGroup)
	NewPriorityRouter().Load(resourceGroup)
	NewTypeRouter().Load(resourceGroup)
}
