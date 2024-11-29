package app

import (
	"fmt"
	"jungotest/internal/app/endpoint"
	"jungotest/internal/app/mw"
	"jungotest/internal/app/service"
	"log"

	"github.com/labstack/echo"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	// middleware, проверяющий User-Role на admin
	// вариант 1
	// глобальный middlware, на все ручки
	a.echo.Use(mw.RoleCheck)
	a.echo.GET("/status", a.e.Status)

	// вариант 2
	// только для /status
	//s.GET("/status", Handler, mw.RoleCheck)

	return a, nil

}

func (a *App) Run() error {
	fmt.Println("server started")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
