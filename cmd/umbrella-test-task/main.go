package main

import (
	"jungotest/internal/pkg/app"
	"log"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// Handler заменится на метод Status() структуры Endpoint
//func Handler(ctx echo.Context) error {
//d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
//dur := time.Until(d)
//s := fmt.Sprintf("Количество дней до 01.01.2025: %d", int(dur.Hours()/24))
//
//err := ctx.String(http.StatusOK, s)
//if err != nil {
//return err
//}
//
//return nil
//}

// заменится на RoleCheck()
//func MW(next echo.HandlerFunc) echo.HandlerFunc {
//return func(ctx echo.Context) error {
//val := ctx.Request().Header.Get("User-Role")

//if val == "admin" {
//log.Println("Red button user detected")
//}

//err := next(ctx)
//if err != nil {
//return err
//}

//return nil
//}
//}
