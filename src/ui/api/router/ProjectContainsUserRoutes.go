package router

import (
	"github.com/labstack/echo/v4"
	"rms-backend/src/ui/api/dicontainer"
	"rms-backend/src/ui/api/handlers"
)

type projectContainsUserRouter struct {
	projectContainsUserHandler handlers.ProjectContainsUserHandlers
}

func NewProjectContainsUserRouter() Router {
	projectContainsUserServices := dicontainer.ProjectContainsUserServices()
	userServices := dicontainer.UserServices()
	projectContainsUserHandlers := handlers.NewProjectContainsUserHandlers(projectContainsUserServices, userServices)
	return &projectContainsUserRouter{projectContainsUserHandler: projectContainsUserHandlers}
}

func (this *projectContainsUserRouter) Load(apiGroup *echo.Group) {
	projectContainsUserGroup := apiGroup.Group("/collaborators")
	projectContainsUserGroup.POST("/:projectID", this.projectContainsUserHandler.Create)
	projectContainsUserGroup.GET("/:projectID", this.projectContainsUserHandler.List)
	projectContainsUserGroup.DELETE("/:projectID/:userID", this.projectContainsUserHandler.Delete)
}
