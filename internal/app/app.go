package app

import (
	"github.com/labstack/echo/v4"
	"github.com/sk10az/go_anime_crud/internal/app/service"
	"github.com/sk10az/go_anime_crud/internal/http/endpoint"
	"github.com/sk10az/go_anime_crud/pkg/logger"
)

type App struct {
	logger   logger.Interface
	service  service.Interface
	endpoint *endpoint.Endpoint
	echo     *echo.Echo
}

func New() (*App, error) {
	var err error
	a := &App{}

	a.logger = logger.New("Info")
	a.logger.Info("Binding logger to application")

	a.service, err = service.New()
	if err != nil {
		a.logger.Error("Error while setting up the services")
		return nil, err
	}

	a.endpoint = endpoint.New(a.service, a.logger)

	a.logger.Info("Setting up Echo instance")
	a.echo = echo.New()

	a.echo.GET("/", a.endpoint.Status)

	return a, nil
}

func (a *App) Run() error {
	a.logger.Info("Running server")

	err := a.echo.Start(":8080")
	if err != nil {
		a.logger.Fatal(err)
	}

	return nil
}