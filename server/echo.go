package server

import (
	"context"
	"github.com/labstack/echo/v4/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VersionInfo struct {
	Version string
	Commit  string
	Build   string
}

type Server struct {
	VersionInfo
	port string
}

func New(info VersionInfo, port string) *Server {
	return &Server{
		info,
		port,
	}
}

func (s Server) initHandlers(e *echo.Echo) {
	e.GET("/", handler)

	e.Any("/*", func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	})
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World, from echo router\n")
}

func (s Server) Serve(ctx context.Context) error {
	e := echo.New()
	e.HideBanner= true
	e.Use(middleware.Recover())
	e.Use(middleware.Recover())
	s.initHandlers(e)

	go func() {
		e.Logger.Infof("start server on port: %s", s.port)
		err := e.Start(":"+s.port)
		if err != nil {
			e.Logger.Errorf("start server error: %v", err)
		}
	}()

	<-ctx.Done()
	return e.Shutdown(ctx)
}
