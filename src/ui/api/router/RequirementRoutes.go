package router

import (
	"github.com/labstack/echo/v4"
	"rms-backend/src/ui/api/dicontainer"
	"rms-backend/src/ui/api/handlers"
)

type requirementRouter struct {
	requirementHandler handlers.RequirementHandlers
}

func NewRequirementRouter() Router {
	requirementServices := dicontainer.RequirementServices()
	requirementHandlers := handlers.NewRequirementHandler(requirementServices)
	return &requirementRouter{requirementHandler: requirementHandlers}
}

func (this *requirementRouter) Load(apiGroup *echo.Group) {
	requirementGroup := apiGroup.Group("/requirements")
	requirementGroup.POST("", this.requirementHandler.Create)
	requirementGroup.GET("", this.requirementHandler.List)
	requirementGroup.GET("/:id", this.requirementHandler.Get)
	requirementGroup.PUT("/:id", this.requirementHandler.Update)
	requirementGroup.DELETE("/:id", this.requirementHandler.Delete)
}
