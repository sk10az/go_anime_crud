package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sk10az/go_anime_crud/internal/app/model"
	"github.com/sk10az/go_anime_crud/internal/app/service"
	"github.com/sk10az/go_anime_crud/pkg/logger"
	"net/http"
	"strconv"
)

type Endpoint struct {
	s service.Interface
	l logger.Interface
}

func New(s service.Interface, l logger.Interface) *Endpoint {
	return &Endpoint{
		s: s,
		l: l,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	req := ctx.Request()
	e.l.Info("Got a request: ", req.Header)

	data := e.s.Ping()
	s := fmt.Sprintf("Server data is: %s", data)

	e.l.Info("Sending data to user")
	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}

// POST /api/ac/
func (e *Endpoint) CreateAnimeCharacter(ctx echo.Context) error {
	ac := new(model.AnimeCharacter)
	if err := ctx.Bind(ac); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	e.s.CreateAnimeCharacter(ac)
	return ctx.JSON(http.StatusCreated, ac)
}

// GET /api/ac/:id
func (e *Endpoint) GetAnimeCharacter(ctx echo.Context) error {
	str_id := ctx.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ac, err := e.s.GetAnimeCharacter(model.IdAC(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, ac)
}

// GET /api/ac
func (e *Endpoint) GetAllAnimeCharacters(ctx echo.Context) error {
	slcAC := e.s.GetAllAnimeCharacters()
	return ctx.JSON(http.StatusOK, slcAC)
}

// PUT /api/ac/:id
func (e *Endpoint) UpdateAnimeCharacter(ctx echo.Context) error {
	str_id := ctx.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ac := new(model.AnimeCharacter)
	if err := ctx.Bind(ac); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ac, err = e.s.UpdateAnimeCharacter(model.IdAC(id), ac)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, ac)
}

// DELETE /api/ac/:id
func (e *Endpoint) DeleteAnimeCharacter(ctx echo.Context) error {
	str_id := ctx.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = e.s.DeleteAnimeCharacter(model.IdAC(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}
