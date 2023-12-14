package router

import (
	"github.com/labstack/echo/v4"
	"rms-backend/src/ui/api/dicontainer"
	"rms-backend/src/ui/api/handlers"
)

type authRouter struct {
	authHandler handlers.AuthHandlers
}

func NewAuthRouter() Router {
	authServices := dicontainer.AuthServices()
	authHandlers := handlers.NewAuthHandler(authServices)
	return &authRouter{authHandler: authHandlers}
}

func (this *authRouter) Load(apiGroup *echo.Group) {
	authGroup := apiGroup.Group("/auth")
	authGroup.POST("/login", this.authHandler.Login)
	authGroup.POST("/sign-up", this.authHandler.SignUp)
}
