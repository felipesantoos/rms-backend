package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os"
)

type Router interface {
	Load(*echo.Group)
}

type router struct{}

func New() Router {
	return &router{}
}

func (this *router) Load(group *echo.Group) {
	if os.Getenv("SERVER_MODE") == "dev" || os.Getenv("SERVER_MODE") == "stage" {
		this.LoadDocs(group)
	}
	NewProjectRouter().Load(group)
	NewRequirementRouter().Load(group)
	NewResourceRouter().Load(group)
}

func (this *router) LoadDocs(group *echo.Group) {
	group.GET("/docs/*", echoSwagger.WrapHandler)
	group.GET("/docs", func(c echo.Context) error {
		return c.Redirect(301, "/api/docs/index.html")
	})
}
