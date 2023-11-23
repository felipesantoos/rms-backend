package router

import (
	"github.com/labstack/echo/v4"
	"rms-backend/src/ui/api/dicontainer"
	"rms-backend/src/ui/api/handlers"
)

type originRouter struct {
	originHandler handlers.OriginHandlers
}

func NewOriginRouter() Router {
	originServices := dicontainer.OriginServices()
	originHandlers := handlers.NewOriginHandler(originServices)
	return &originRouter{originHandler: originHandlers}
}

func (this *originRouter) Load(apiGroup *echo.Group) {
	originGroup := apiGroup.Group("/origins")
	originGroup.GET("", this.originHandler.List)
}
