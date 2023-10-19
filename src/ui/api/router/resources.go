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
	projectHandler := handlers.NewResourcesHandler(projectServices)
	return &projectRouter{projectHandler: projectHandler}
}

func (this *projectRouter) Load(apiGroup *echo.Group) {
	projectGroup := apiGroup.Group("/projects")
	projectGroup.POST("", this.projectHandler.Create)
	projectGroup.GET("", this.projectHandler.List)
	projectGroup.GET("/:id", this.projectHandler.Get)
}
