package router

import (
	"github.com/labstack/echo/v4"
	"rms-backend/src/ui/api/dicontainer"
	"rms-backend/src/ui/api/handlers"
)

type typeRouter struct {
	typeHandler handlers.TypeHandlers
}

func NewTypeRouter() Router {
	typeServices := dicontainer.TypeServices()
	typeHandlers := handlers.NewTypeHandler(typeServices)
	return &typeRouter{typeHandler: typeHandlers}
}

func (this *typeRouter) Load(apiGroup *echo.Group) {
	typeGroup := apiGroup.Group("/types")
	typeGroup.GET("", this.typeHandler.List)
}
