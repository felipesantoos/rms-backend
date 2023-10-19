package router

import (
	"github.com/labstack/echo/v4"
	"rms-backend/src/ui/api/dicontainer"
	"rms-backend/src/ui/api/handlers"
)

type projectRouter struct {
	projectHandler handlers.ProjectHandlers
}

func NewProjectRouter() Router {
	projectServices := dicontainer.ProjectServices()
	projectHandlers := handlers.NewProjectHandler(projectServices)
	return &projectRouter{projectHandler: projectHandlers}
}

func (this *projectRouter) Load(apiGroup *echo.Group) {
	projectGroup := apiGroup.Group("/projects")
	projectGroup.POST("", this.projectHandler.Create)
	projectGroup.GET("", this.projectHandler.List)
	projectGroup.GET("/:id", this.projectHandler.Get)
	projectGroup.PUT("/:id", this.projectHandler.Update)
	projectGroup.DELETE("/:id", this.projectHandler.Delete)
}
