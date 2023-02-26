package app

import (
	"github.com/labstack/echo/v4"
	"github.com/sk10az/go_anime_crud/internal/app/repository/memory"
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

	// Setup logger
	a.logger = logger.New("Info")
	a.logger.Info("Binding logger to application")

	// Setup service
	repo, err := memory.New(a.logger)
	if err != nil {
		return nil, err
	}

	a.service, err = service.New(a.logger, repo)
	if err != nil {
		a.logger.Error("Error while setting up the services")
		return nil, err
	}

	// Setup server side handler
	a.endpoint = endpoint.New(a.service, a.logger)

	a.logger.Info("Setting up the Echo instance")
	a.echo = echo.New()

	a.echo.GET("/", a.endpoint.Status)

	a.echo.POST("/api/ac", a.endpoint.CreateAnimeCharacter)
	a.echo.GET("/api/ac/:id", a.endpoint.GetAnimeCharacter)
	a.echo.GET("/api/ac/", a.endpoint.GetAllAnimeCharacters)
	a.echo.PUT("/api/ac/:id", a.endpoint.UpdateAnimeCharacter)
	a.echo.DELETE("/api/ac/:id", a.endpoint.DeleteAnimeCharacter)

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
