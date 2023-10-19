package router

import (
	"github.com/labstack/echo/v4"
	"rms-backend/src/ui/api/dicontainer"
	"rms-backend/src/ui/api/handlers"
)

type priorityRouter struct {
	priorityHandler handlers.PriorityHandlers
}

func NewPriorityRouter() Router {
	priorityServices := dicontainer.PriorityServices()
	priorityHandlers := handlers.NewPriorityHandler(priorityServices)
	return &priorityRouter{priorityHandler: priorityHandlers}
}

func (this *priorityRouter) Load(apiGroup *echo.Group) {
	priorityGroup := apiGroup.Group("/priorities")
	priorityGroup.GET("", this.priorityHandler.List)
}
