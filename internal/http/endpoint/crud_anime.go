package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sk10az/go_anime_crud/internal/app/service"
	"github.com/sk10az/go_anime_crud/pkg/logger"
	"net/http"
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

	data := e.s.Data()
	s := fmt.Sprintf("Server data is: %s", data)

	e.l.Info("Sending data to user")
	err := ctx.String(http.StatusOK, s);
	if err != nil {
		return err
	}
	return nil
}
