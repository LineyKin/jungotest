package endpoint

// endpoint - первый слой, он взаимодействует с клиентом и возвращает ему какой-то ответ
// сервисный слой делает непосредственно бизнес-логику

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Service interface {
	DaysLeft() int
}

// типа класс
type Endpoint struct {
	s Service
}

// типа конструктор класса
func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.s.DaysLeft()
	s := fmt.Sprintf("Количество дней до 01.01.2025: %d", d)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}
