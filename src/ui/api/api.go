package api

import (
	"fmt"
	"rms-backend/src/core"
	_ "rms-backend/src/ui/api/docs"
	"rms-backend/src/ui/api/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

var logger = core.CoreLogger().With().Logger()

type API interface {
	Serve()
}

type api struct {
	host   string
	port   int
	server *echo.Echo
}

// NewAPI
// @title RMS API
// @version 1.0
// @description RMS backend API
// @contact.name Felipe da Silva Santos
// @contact.email fss30@aluno.ifal.edu.br
// @BasePath /api
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func NewAPI(host string, port int) API {
	server := echo.New()
	return &api{host, port, server}
}

func (a *api) Serve() {
	a.setupMiddlewares()
	a.loadRoutes()
	a.start()
}

func (a *api) setupMiddlewares() {
	a.server.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:   true,
		LogError:    true,
		LogRemoteIP: true,
		LogURIPath:  true,
		LogURI:      true,
		LogStatus:   true,
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			var event *zerolog.Event
			if v.Status < 400 {
				event = logger.Info()
			} else if v.Status >= 400 {
				event = logger.Error()
			}
			event.Str("PATH", v.URIPath).Str("REMOTEIP", v.RemoteIP).Int("STATUS", v.Status)
			if v.Error != nil {
				event.Str("ERROR", v.Error.Error())
			}
			event.Msg(fmt.Sprintf("[%-5s]", v.Method))
			return nil
		},
	}))
	a.server.Use(middleware.Recover())
}

func (a *api) rootGroup() *echo.Group {
	return a.server.Group("/api")
}

func (a *api) loadRoutes() {
	router := router.New()
	router.Load(a.rootGroup())
}

func (a *api) start() {
	address := fmt.Sprintf("%s:%d", a.host, a.port)
	err := a.server.Start(address)
	a.server.Logger.Fatal(err)
}
